package page

/**
    INT8		INT			INT			LONG				LONG
  INDEX # PAGE_NUMBER # PAGE_SIZE # LOWEST_TIMESTAMP # HIGHEST_TIMESTAMP #
*/
type PageHeader struct {
	// Used for navigation
	PageNumber uint16
	// Read from the Index Map
	KeyIndex int8
	// Max Amount of PageSize is 65536 elements, means it can contain
	// 65536 PageCell objects
	PageSize         uint16
	LowestTimeStamp  int
	HighestTimeStamp int
}
