package page

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
	binary.Write(buffer, binary.BigEndian, ph.PageNumber)
	binary.Write(buffer, binary.BigEndian, ph.KeyIndex)
	binary.Write(buffer, binary.BigEndian, ph.PageSize)
	binary.Write(buffer, binary.BigEndian, ph.LowestTimeStamp)
	binary.Write(buffer, binary.BigEndian, ph.HighestTimeStamp)

	return buffer.Bytes()
}

func (ps *PageSerializer) serializeCell(pc *entity.PageCell) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, pc.DataType)
	binary.Write(buffer, binary.BigEndian, pc.Content)
	binary.Write(buffer, binary.BigEndian, pc.Label)
	binary.Write(buffer, binary.BigEndian, pc.Key)
	binary.Write(buffer, binary.BigEndian, pc.Offset)

	return buffer.Bytes()
}

func (ps *PageSerializer) SerializePage(page *entity.Page) []byte {
	// Concat of header+cell
	header := ps.serializeHeader(&page.Header)
	cells := make([]byte, len(page.Cells)*49)

	for _, c := range page.Cells {
		cells = append(cells, ps.serializeCell(&c)...)
	}

	return append(header, cells...)
}

func (ps *PageSerializer) deserializeHeader(data []byte) *entity.PageHeader {
	platformIntSize := unsafe.Sizeof(0)
	pageNumber := binary.BigEndian.Uint16(data[0:1])                                             // 0...1 byte
	keyIndex := binary.BigEndian.Uint16(data[1:2])                                               // +1 byte
	pageSize := binary.BigEndian.Uint16(data[2:4])                                               // +2 bytes
	lowestTimeStamp := binary.BigEndian.Uint32(data[4 : 4+platformIntSize])                      // + int bytes
	highestTimeStamp := binary.BigEndian.Uint16(data[4+platformIntSize : 4+(2*platformIntSize)]) // + 4 bytes

	return &entity.PageHeader{
		PageNumber:       int8(pageNumber),
		KeyIndex:         int8(keyIndex),
		PageSize:         pageSize,
		LowestTimeStamp:  int(lowestTimeStamp),
		HighestTimeStamp: int(highestTimeStamp),
	}

}

func (ps *PageSerializer) deserializeCell(data []byte) *entity.PageCell {

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
	header := ps.deserializeHeader(*bytes)
	// todo check how many cells are in header and then from it deserialize
	// the cells

	return &entity.Page{
		Header: *header,
	}
}
