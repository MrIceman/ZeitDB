package entity

type Page struct {
	Header *PageHeader
	Cells  *[]PageCell
}

/**
    INT8		INT			INT			LONG				LONG
  INDEX # PAGE_NUMBER # PAGE_SIZE # LOWEST_TIMESTAMP # HIGHEST_TIMESTAMP #

  PAGE_NUMBER | KEY_INDEX | PAGE_SIZE | LOWEST_TIMESTAMP | HIGHEST_TIMESTAMP
	1 Byte		1 Byte		2 Bytes		4 Bytes				4 Bytes

	12 Bytes in total
*/
type PageHeader struct {
	// Used for navigation
	PageNumber int8
	// Read from the Index Map
	KeyIndex int8
	// Max Amount of PageSize is 65536 elements, means it can contain
	// 65536 PageCell objects
	PageSize         uint16
	LowestTimeStamp  int64
	HighestTimeStamp int64
	IndexFileName    string
}

/**
  A page cell has following structure
	KEY | LABEL | DATA_TYPE | LENGTH | CONTENT

	A label is a 4 byte integer and is mapped within a label-table to the
	full qualifier
*/
type PageCell struct {
	/**
	  0x01  = int
	  0x02  = string
	  0x03  = bool
	*/
	Key      int64
	Label    int32
	DataType byte
	Length   int32
	Content  string
}
