package io

import (
	"ZeitDB/entity"
	"bytes"
	"encoding/binary"
	"unsafe"
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

	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func (ps *PageSerializer) serializeCell(pc *entity.PageCell) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, pc.DataType)
	err = binary.Write(buffer, binary.BigEndian, pc.Content)
	err = binary.Write(buffer, binary.BigEndian, pc.Label)
	err = binary.Write(buffer, binary.BigEndian, pc.Key)
	err = binary.Write(buffer, binary.BigEndian, pc.Offset)

	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
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

func (ps *PageSerializer) DeserializeHeader(data []byte) *entity.PageHeader {
	int32Size := unsafe.Sizeof(int32(0))
	pageNumber := binary.BigEndian.Uint16(data[0:1])                                 // 0...1 byte
	keyIndex := binary.BigEndian.Uint16(data[1:2])                                   // +1 byte
	pageSize := binary.BigEndian.Uint16(data[2:4])                                   // +2 bytes
	lowestTimeStamp := binary.BigEndian.Uint32(data[4 : 4+int32Size])                // + int bytes
	highestTimeStamp := binary.BigEndian.Uint32(data[4+int32Size : 4+(2*int32Size)]) // + 4 bytes

	return &entity.PageHeader{
		PageNumber:       int8(pageNumber),
		KeyIndex:         int8(keyIndex),
		PageSize:         pageSize,
		LowestTimeStamp:  int32(lowestTimeStamp),
		HighestTimeStamp: int32(highestTimeStamp),
	}

}

func (ps *PageSerializer) DeserializeCell(data []byte) *entity.PageCell {

	dataType := data[0]                            // 1st byte
	content := string(data[1:17])                  // +8 bytes
	label := string(data[17:33])                   // +16 bytes
	key := binary.BigEndian.Uint32(data[33:37])    // + 4 bytes
	offset := binary.BigEndian.Uint16(data[37:39]) // + 2 bytes

	return &entity.PageCell{
		DataType: dataType,
		Content:  content,
		Label:    label,
		Key:      key,
		Offset:   offset,
	}
}

func (ps *PageSerializer) DeserializePage(bytes *[]byte) *entity.Page {
	header := ps.DeserializeHeader(*bytes)
	// todo check how many cells are in header and then from it deserialize
	// the cells

	return &entity.Page{
		Header: header,
	}
}
