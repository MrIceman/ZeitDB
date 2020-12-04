package io

import (
	"ZeitDB/entity"
	"testing"
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

	if len(bytes) != 12 {
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
		t.Error("Invalid LowestTimeStamp")
	}
	if serializedHeader.HighestTimeStamp != 1780 {
		t.Error("Invalid HighestTimeStamp")
	}
}
