package page

import (
	"ZeitDB/entity"
	"bytes"
	"encoding/binary"
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

func (ps *PageSerializer) deserializeHeader(bytes *[]byte) *entity.PageHeader {

}

func (ps *PageSerializer) deserializeCell(bytes *[]byte) *entity.PageCell {

}

func (ps *PageSerializer) DeserializePage(bytes *[]byte) *entity.Page {

}
