package page

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
	Offset   int
}
