package entity

type Page struct {
	Header PageHeader
	Cells  []PageCell
}

/**
    INT8		INT			INT			LONG				LONG
  INDEX # PAGE_NUMBER # PAGE_SIZE # LOWEST_TIMESTAMP # HIGHEST_TIMESTAMP #
*/
type PageHeader struct {
	// Used for navigation
	PageNumber int8
	// Read from the Index Map
	KeyIndex int8
	// Max Amount of PageSize is 65536 elements, means it can contain
	// 65536 PageCell objects
	PageSize         uint16
	LowestTimeStamp  int
	HighestTimeStamp int
}

/**
  A page cell has a fixed size of
   1 + 16 + 16 + 16 = 49 bytes
*/
type PageCell struct {
	/**
	  0x01  = int
	  0x02  = string
	  0x03  = bool
	*/
	DataType byte
	Content  string
	Label    string
	Key      uint32
	Offset   uint16
}
