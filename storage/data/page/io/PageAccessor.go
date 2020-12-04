package io

import (
	"ZeitDB/entity"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"unsafe"
)

func ReadPage(
	path string,
	serializer *PageSerializer) (*entity.Page, error) {
	println("path: ", path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		pageHeaderRaw := make([]byte, unsafe.Sizeof(entity.PageHeader{}))
		// We need to readLen the PageHeader first to evaluate
		// how many cells have to be readLen

		println("Size of pageHeaderRaw ", pageHeaderRaw)
		readLen, err := f.ReadAt(pageHeaderRaw, 0)
		if readLen != len(pageHeaderRaw) {
			return nil, errors.New("invalid page length. Expected " + strconv.Itoa(len(pageHeaderRaw)) + "but received" + strconv.Itoa(readLen))
		} else if err != nil {
			panic(err)
		}

	} else {
		panic(err)
	}
	return nil, nil
}

func WritePage(path string, page *entity.Page, serializer *PageSerializer) error {
	bytes := serializer.SerializePage(page)
	println("Writing ", unsafe.Sizeof(*page.Header), " bytes ", unsafe.Sizeof(*page))
	err := ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
