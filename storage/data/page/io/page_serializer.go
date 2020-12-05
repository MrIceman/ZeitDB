package io

import (
	"ZeitDB/entity"
	"bytes"
	"encoding/binary"
)

type PageSerializer struct {
}

func (ps *PageSerializer) serializeHeader(ph *entity.PageHeader) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, ph.PageNumber)
	err = binary.Write(buffer, binary.BigEndian, ph.KeyIndex)
	err = binary.Write(buffer, binary.BigEndian, ph.PageSize)
	err = binary.Write(buffer, binary.BigEndian, ph.LowestTimeStamp)
	err = binary.Write(buffer, binary.BigEndian, ph.HighestTimeStamp)
	err = binary.Write(buffer, binary.BigEndian, []byte(ph.IndexFileName))

	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func (ps *PageSerializer) DeserializeHeader(data []byte) *entity.PageHeader {
	intBytes := 8
	pageNumber := data[0]                                                          // 1 byte
	keyIndex := data[1]                                                            // 1 byte
	pageSize := binary.BigEndian.Uint16(data[2:4])                                 // 2 bytes
	lowestTimeStamp := binary.BigEndian.Uint64(data[4 : 4+intBytes])               // 4 bytes
	highestTimeStamp := binary.BigEndian.Uint64(data[4+intBytes : 4+(2*intBytes)]) // 4 bytes

	return &entity.PageHeader{
		PageNumber:       int8(pageNumber),
		KeyIndex:         int8(keyIndex),
		PageSize:         pageSize,
		LowestTimeStamp:  int64(lowestTimeStamp),
		HighestTimeStamp: int64(highestTimeStamp),
	}

}

func (ps *PageSerializer) serializeCell(pc *entity.PageCell) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, pc.Key)
	err = binary.Write(buffer, binary.BigEndian, pc.Label)
	err = binary.Write(buffer, binary.BigEndian, pc.DataType)
	err = binary.Write(buffer, binary.BigEndian, pc.Length)
	err = binary.Write(buffer, binary.BigEndian, []byte(pc.Content))

	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func (ps *PageSerializer) deserializeCell(data []byte) *entity.PageCell {
	key := binary.BigEndian.Uint64(data[0:8])
	label := binary.BigEndian.Uint32(data[8:12])
	dataType := data[12] // 1 byte
	length := binary.BigEndian.Uint32(data[13:17])
	content := string(data[17 : 17+length])

	return &entity.PageCell{
		DataType: dataType,
		Content:  content,
		Label:    int32(label),
		Key:      int64(key),
		Length:   int32(length),
	}
}

func (ps *PageSerializer) SerializePage(page *entity.Page) []byte {
	// Concat of header+cell
	header := ps.serializeHeader(page.Header)
	println("PageHeader is ", len(header), " bytes long")
	cells := make([]byte, len(*page.Cells)*49)

	for _, c := range *page.Cells {
		cells = append(cells, ps.serializeCell(&c)...)
	}

	return append(header, cells...)
}

func (ps *PageSerializer) DeserializePage(bytes *[]byte) *entity.Page {
	header := ps.DeserializeHeader(*bytes)
	// todo check how many cells are in header and then from it deserialize
	// the cells

	return &entity.Page{
		Header: header,
	}
}
