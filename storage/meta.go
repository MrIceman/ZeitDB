package storage

import (
	"bytes"
	"encoding/binary"
	"os"
)

type MetaInfo struct {
	GlobalLowestTimeStamp  uint32
	GlobalHighestTimeStamp uint32
	// GlobalLowestTimeStamp is the first cell of the first page
	// GlobalHighestTimeStamp is the last cell of the last page
	AmountOfPages uint32
	AmountOfCells uint32
	Version       uint32
}

func (m *MetaInfo) ToByteArray() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, m.GlobalLowestTimeStamp)
	binary.Write(buffer, binary.BigEndian, m.GlobalHighestTimeStamp)
	binary.Write(buffer, binary.BigEndian, m.AmountOfPages)
	binary.Write(buffer, binary.BigEndian, m.AmountOfCells)
	binary.Write(buffer, binary.BigEndian, m.Version)
	return buffer.Bytes()
}

func FromByteArray(data []byte) *MetaInfo {
	byteSize := 4
	var gl = binary.BigEndian.Uint32(data[0:byteSize])
	var gh = binary.BigEndian.Uint32(data[byteSize : byteSize*2])
	var aop = binary.BigEndian.Uint32(data[byteSize*2 : byteSize*3])
	var aoc = binary.BigEndian.Uint32(data[byteSize*3 : byteSize*4])
	var ver = binary.BigEndian.Uint32(data[byteSize*4 : byteSize*5])

	return &MetaInfo{
		GlobalHighestTimeStamp: gh,
		GlobalLowestTimeStamp:  gl,
		AmountOfCells:          aoc,
		AmountOfPages:          aop,
		Version:                ver,
	}
}

type Configuration struct {
	MetaInfoFilePath string
}

func (c *Configuration) ObtainMetaInfo() (*MetaInfo, error) {
	f, err := os.Open(c.MetaInfoFilePath)
	if err != nil {
		defer f.Close()
		byteArray := make([]byte, 20)
		_, err = f.Read(byteArray)
		if err != nil {
			return FromByteArray(byteArray), nil
		}
		return nil, err
	}
	return nil, err
}
