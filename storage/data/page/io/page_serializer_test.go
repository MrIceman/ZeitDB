package io

import (
	"ZeitDB/entity"
	"testing"
	"time"
)

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
	content := "1jiyvasiofjiojiosjfiosjfiojasoifjoisajiofjiojiosjfiosjfiojasoifjoisajiofjiojiosjfiosjfiojasoifjoisajiofjiojiosjfiosjfiojasoifjoisajiofjiojiosjfiosjfiojasoifjoisajiof"
	cell := entity.PageCell{
		DataType: 0x9,
		Content:  content,
		Label:    0,
		Key:      time.Now().Unix(),
	}
	serializer := PageSerializer{}
	cellMemorySize := 8 + 4 + 1 + 4 + len([]byte(content))
	serializedCell := serializer.serializeCell(&cell)
	if len(serializedCell) != cellMemorySize {
		t.Error("Serialized Cell is not ", cellMemorySize, " bytes long, but ", len(serializedCell))
	}
	//deserializedCell := serializer.deserializeCell(serializedCell)

}

func TestPageSerializer_DeserializePage(t *testing.T) {
	header := entity.PageHeader{
		PageNumber:       10,
		KeyIndex:         4,
		PageSize:         13,
		LowestTimeStamp:  10,
		HighestTimeStamp: 1780,
	}

	p := entity.Page{
		Header: &header,
		Cells: &[]entity.PageCell{
			{
				DataType: 20,
				Content:  "1jiyvasiof sajiof",
				Label:    3,
				Key:      20,
			},
		},
	}
	serializer := PageSerializer{}
	serializedPage := serializer.SerializePage(&p)
	deserializedPage := serializer.DeserializePage(&serializedPage)
	println(deserializedPage)
}
