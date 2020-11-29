package entity

import (
	"testing"
	"unsafe"
)

func TestDeserialization(t *testing.T) {
	metaInfo := MetaInfo{
		GlobalLowestTimeStamp:  7980,
		GlobalHighestTimeStamp: 8,
		AmountOfPages:          1,
		AmountOfCells:          8,
		Version:                1,
	}
	println("size of integer:", unsafe.Sizeof(int(8)))

	serialized := metaInfo.ToByteArray()
	deserialized := FromByteArray(serialized)

	println("globallowesttimestamp:", deserialized.GlobalLowestTimeStamp)
	if deserialized.Version != 1 {
		t.Error("Version is not 1")
	}
	if deserialized.AmountOfPages != 1 {
		t.Error("AmountOfPages is not 1")
	}
}
