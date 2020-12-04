package entity

import (
	"bytes"
	"encoding/binary"
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
	err := binary.Write(buffer, binary.BigEndian, m.GlobalLowestTimeStamp)
	err = binary.Write(buffer, binary.BigEndian, m.GlobalHighestTimeStamp)
	err = binary.Write(buffer, binary.BigEndian, m.AmountOfPages)
	err = binary.Write(buffer, binary.BigEndian, m.AmountOfCells)
	err = binary.Write(buffer, binary.BigEndian, m.Version)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func FromByteArray(data []byte) *MetaInfo {
	byteSize := 4
	var globalLowestTs = binary.BigEndian.Uint32(data[0:byteSize])
	var globalHighestTs = binary.BigEndian.Uint32(data[byteSize : byteSize*2])
	var amountOfpages = binary.BigEndian.Uint32(data[byteSize*2 : byteSize*3])
	var amountOfCells = binary.BigEndian.Uint32(data[byteSize*3 : byteSize*4])
	var version = binary.BigEndian.Uint32(data[byteSize*4 : byteSize*5])

	return &MetaInfo{
		GlobalHighestTimeStamp: globalHighestTs,
		GlobalLowestTimeStamp:  globalLowestTs,
		AmountOfCells:          amountOfCells,
		AmountOfPages:          amountOfpages,
		Version:                version,
	}
}
