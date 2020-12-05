package io

import (
	"ZeitDB/entity"
	"testing"
)

func TestSerializeIndexPage(t *testing.T) {
	indexPage := entity.IndexPage{
		PageNumber: 10,
		Entries:    nil,
	}

	serializer := IndexSerializer{}

	serializedPage := serializer.SerializeIndexPage(indexPage)

	if len(serializedPage) != 8 {
		t.Error("Serialized Page is not 8 bytes long")
	}
}
