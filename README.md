# ZeitDB
ZeitDB is a time-series, push based lightweight disk database. The main aim currently is to learn more about database</br>
internals, practice Go-Lang skills and play around with latest research regarding data storage. <br/>
Right now it is <font color="red"><b>heavily Under Construction </b></font>
<br/>
## Architectural draft
Transport -> Query Processor -> Execution Engine -> Storage Engine

## How does the Page Layout work
See code

# Insertion
Every entry has a timestamp. The storage engine receives the timestamp, reads the page dictionary to find<br/>
the page that contains the timestamp. This happens through a binary search. Once the page is found, the algorithm </br>
performs a binary search to find the proper position for our entry to be inserted. Compared with our inserted node,<br/>
its predecessor is the biggest smallest node, and the successor is the lowest highest node. Eventually, a page <br/>
overflow might happen. See the example below for more information 


<br/><h4>Example: </h4>
<b>Given Timestamps:</b> t0, t1, t2<br/><br/>t1 is the timestamp we want to insert, t0 is the lowest timestamp
recorded until now (lowest boundary) and t2 is the highest  one (highest boundary), so that t0 < t1 < t2.<br/> The algorithm would then 
read the page dictionary and try to find the right page number for t1.<br/>Let's say we have 20 pages, and
t0 is on the 1st page and t2 on the 20th, then the algorithm would grab
the 10th page, check the page header and see if 
<br/><code>page_lowest_timestamp <= t1 <= page_highest_timestamp</code><br/> where the page_lowest_timestamp and page_highest_timestamp are 
the timestamp extremas (boundaries) contained within the page. <br/> The algorithm would proceed with a binary search 
until a page is found with <b>t1</b> between its boundaries. Once the page is found, we
If the timestamp is contained, then we're
performing the same binary search within the page. For that, we'e reading the PageSize and then navigate to
the entry PageSize/2 and compare the values. We continue until we find the right place where
t1 is higher than each timestap on the left side and lower than every timestamp on the right side.
After insertion we're checking if the page has overflown, means we have more than
2^16 entries. If this is the case, we need to create a new page and
every following page gets its page number incremented (we're shifting the pages).

# Lookup
A lookup happens similary to the Insertion, via binary search.