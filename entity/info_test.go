package entity

import (
	"testing"
	"time"
)

func TestDeserialization(t *testing.T) {
	metaInfo := MetaInfo{
		GlobalLowestTimeStamp:  7980,
		GlobalHighestTimeStamp: 8,
		AmountOfPages:          100,
		AmountOfCells:          8,
		Version:                165,
	}

	serialized := metaInfo.ToByteArray()
	deserialized := FromByteArray(serialized)

	if deserialized.Version != 165 {
		t.Error("Version is not 1")
	}
	if deserialized.AmountOfPages != 100 {
		t.Error("AmountOfPages is not 100")
	}

	startTime := time.Now()
	for i := 0; i < 10000000; i++ {
		metaInfo.Version = uint32(i) + metaInfo.Version
		serialized = metaInfo.ToByteArray()
		deserialized = FromByteArray(serialized)
	}

	totalTime := time.Now().Sub(startTime)

	println("Took ", totalTime.Milliseconds(), " Milliseconds")
}
