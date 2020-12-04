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
		HighestTimeStamp: 10,
	}

	serializer := PageSerializer{}

	bytes := serializer.serializeHeader(&header)

	if len(bytes) != 12 {
		t.Error("Serialized Headers is not 20 bytes long but ", len(bytes))
	}
}
