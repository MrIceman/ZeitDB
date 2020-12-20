package io

import (
	"ZeitDB/entity"
	"testing"
	"time"
	"unsafe"
)

type blah struct {
	asas  int32
	iasas int32
	asas2 int8
}

func TestHeaderLength(t *testing.T) {
	var asfio string
	println("blah blah", asfio != "")
	headerSize := unsafe.Sizeof(entity.PageHeader{})
	println(unsafe.Sizeof(int64(0)))
	println(unsafe.Sizeof(uint16(0)))
	println(unsafe.Sizeof(int8(0)))
	println("header size: ", headerSize)
	if headerSize != 24 {
		t.Error("Header is not 24 bytes long but ", headerSize)
	}
}

func TestPageSerializer_DeserializeHeader(t *testing.T) {
	header := entity.PageHeader{
		PageNumber:       10,
		KeyIndex:         4,
		PageSize:         13,
		LowestTimeStamp:  10,
		HighestTimeStamp: 1780,
	}

	serializer := PageSerializer{}
	bytes := serializer.serializeHeader(&header)

	if len(bytes) != 20 {
		t.Error("Serialized Headers is not 20 bytes long but ", len(bytes))
	}

	serializedHeader := serializer.DeserializeHeader(bytes)

	if serializedHeader.PageNumber != 10 {
		t.Error("Invalid Page Number")
	}
	if serializedHeader.KeyIndex != 4 {
		t.Error("Invalid Key Index")
	}
	if serializedHeader.PageSize != 13 {
		t.Error("Invalid PageSize")
	}
	if serializedHeader.LowestTimeStamp != 10 {
		t.Error("Invalid LowestTimeStamp: ", serializedHeader.LowestTimeStamp)
	}
	if serializedHeader.HighestTimeStamp != 1780 {
		t.Error("Invalid HighestTimeStamp: ", serializedHeader.HighestTimeStamp)
	}
}

func TestPageSerializer_DeserializeCell(t *testing.T) {
	content := "1jiyvasiof sajiof"
	cell := entity.PageCell{
		DataType: 0x9,
		Content:  content,
		Label:    0,
		Length:   -1,
		Key:      time.Now().Unix(),
	}
	serializer := PageSerializer{}
	cellMemorySize := 8 + 4 + 1 + 4 + len([]byte(content))
	serializedCell := serializer.serializeCell(&cell)
	if len(serializedCell) != cellMemorySize {
		t.Error("Serialized Cell is not ", cellMemorySize, " bytes long, but ", len(serializedCell))
	}
	deserializedCell := serializer.deserializeCell(serializedCell)
	deserializedCellSize := 8 + 4 + 1 + 4 + int(deserializedCell.Length)
	println("Deserialized Cell length:", deserializedCellSize)
	if deserializedCellSize != cellMemorySize {
		t.Error("Deserialized Cell is not ", cellMemorySize, " bytes long, but ", deserializedCellSize)
	}

}

func TestPageSerializer_DeserializePage(t *testing.T) {
	header := entity.PageHeader{
		PageNumber:       10,
		KeyIndex:         4,
		PageSize:         1,
		LowestTimeStamp:  10,
		HighestTimeStamp: 1780,
		Magic:            0,
	}
	cell := entity.PageCell{
		DataType: 20,                           // 1 byte
		Content:  "1jjiojiojioiyvasiof sajiof", // 26 bytes
		Label:    3,                            // 4 bytes
		Key:      20,                           // 8 bytes
	}

	p := entity.Page{
		Header: &header,
		Cells: &[]entity.PageCell{
			cell,
		},
	}
	serializer := PageSerializer{}
	serializedPage := serializer.SerializePage(&p)
	deserializedPage := *serializer.DeserializePage(serializedPage)

	if len(*deserializedPage.Cells) != 1 {
		t.Error("There should be 1 cell deserialized but instead we got", len(*deserializedPage.Cells))
	}

	// assert values
	deserializedCell := (*deserializedPage.Cells)[0]

	if int(deserializedCell.Length) != len([]byte((*p.Cells)[0].Content)) {
		t.Error("Cell content was not deserialized properly. The length does not match.")
	}
	if deserializedCell.DataType != cell.DataType {
		t.Error("DataType is not as expected")
	}
	if deserializedCell.Content != cell.Content {
		t.Error("Cell is not as expected")
	}
	if deserializedCell.Key != cell.Key {
		t.Error("Key is not as expected")
	}
	if deserializedCell.Label != cell.Label {
		t.Error("Label is not as expected")
	}
}
