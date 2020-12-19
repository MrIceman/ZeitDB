package entity

import "strconv"

type Page struct {
	Header *PageHeader
	Cells  *[]PageCell
}

func (p *Page) ToString() string {
	return (*p.Header).ToString() + " / cells: " + strconv.Itoa(len(*p.Cells))
}

/**
    INT8		INT			INT			LONG				LONG
  INDEX # PAGE_NUMBER # PAGE_SIZE # LOWEST_TIMESTAMP # HIGHEST_TIMESTAMP #

  PAGE_NUMBER | KEY_INDEX | PAGE_SIZE | LOWEST_TIMESTAMP | HIGHEST_TIMESTAMP
	1 Byte		1 Byte		2 Bytes		4 Bytes				4 Bytes

	12 Bytes in total
*/
type PageHeader struct {
	// Added a placeholder so the PageHeader is a multiple
	// of 8 (24 in this case), and we can avoid alignment
	Magic int32
	// Used for navigation
	PageNumber int8
	// Read from the Index Map
	KeyIndex int8
	// Max Amount of PageSize is 65536 elements, means it can contain
	// 65536 PageCell objects
	PageSize         uint16
	LowestTimeStamp  int64
	HighestTimeStamp int64
}

func (ph *PageHeader) ToString() string {
	return "hst: " + int64ToString(ph.HighestTimeStamp) + " lst:" + int64ToString(ph.LowestTimeStamp) + " pageSize:" +
		int16ToString(ph.PageSize) + " keyIndex: " + int8ToString(ph.KeyIndex) +
		" pageNumber:" + int8ToString(ph.PageNumber)
}

func int64ToString(anyInt int64) string {
	return strconv.Itoa(int(anyInt))
}

func int32ToString(anyInt int32) string {
	return strconv.Itoa(int(anyInt))
}
func int16ToString(anyInt uint16) string {
	return strconv.Itoa(int(anyInt))
}

func int8ToString(anyInt int8) string {
	return strconv.Itoa(int(anyInt))
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
