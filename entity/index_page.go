package entity

/**
Every page has an IndexPage which holds an index for ever entry
It is the index page that has to be searched, maintained and sorted continuously.
All the Entries within an IndexPage are sorted.
*/
type IndexPage struct {
	PageNumber int8
	Entries    []IndexEntry
}

type IndexEntry struct {
	Key             int64
	offset          int32
	markedForDelete bool
}
