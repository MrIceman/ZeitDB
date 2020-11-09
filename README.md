# ZeitDB
ZeitDB is a time-series, push based lightweight disk-based database.<br/>
<br/>
## Architectural draft
Transport -> Query Processor -> Execution Engine -> Storage Engine

## How does the Page Layout work
See code

# Insertion
Every entry has a timestamp. The storage engine takes the timestamp, reads the meta file
and checks what the lowest timestamp entry is, what the highest is,
how many page entries there are and then performs a binary
search to find the right page.
Let's say we have a timestamp t1,
a timestamp t0 and a timestamp t2, where t1 is the timestamp we want to insert, t0 is the lowest timestamp
recorded until now and t2 is the highest  one. Also t0 < t1 < t2, the algorithm would then 
read the page meta data and try to find the right page number for t1.  If we have 20 pages, where
t0 is on the 1st page (the first one), and t2 on the 20th, then the algorithm would grab
the 10th page, check the page header and see if 
page_lowest_timestamp <= t1 <= page_highest_timestamp, where the page_lowest_timestamp and page_highest_timestamp are 
the timestamps contained within the page. If the timestamp is contained, then we're
performing the same binary search within the page. For that, we'e reading the PageSize and then navigate to
the entry PageSize/2 and compare the values. We continue until we find the right place where
t1 is higher than each timestap on the left side and lower than every timestamp on the right side.
After insertion we're checking if the page has overflown, means we have more than
2^16 entries. If this is the case, we need to create a new page and
every following page gets its page number incremented (we're shifting the pages).

# Lookup
A lookup happens similary to the Insertion, via binary search.