package io

import (
	"ZeitDB/entity"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"unsafe"
)

/**
	The Page is composed of a PageHeader and an array of PageCells.
    |PageHeader| PC 0 | PC 1 | PC 2 | ... | PC 3
*/
func ReadPage(
	path string,
	serializer *PageSerializer) (*entity.Page, error) {
	f, err := os.Open(path)
	println("opening path", path)
	if err == nil {
		defer f.Close()
		pageHeaderRaw := make([]byte, unsafe.Sizeof(entity.PageHeader{}))
		// We need to readLen the PageHeader first to evaluate
		// how many cells have to be readLen

		readLen, err := f.ReadAt(pageHeaderRaw, 0)
		header := serializer.DeserializeHeader(pageHeaderRaw)
		if (readLen) != len(pageHeaderRaw) {
			panic(errors.New("invalid page length. Expected " + strconv.Itoa(len(pageHeaderRaw)) + "but received" + strconv.Itoa(readLen)))
		} else if err != nil {
			panic(err)
		}
		var pageCells []entity.PageCell
		if header.PageSize > 0 {
			// Parse Page Cells
			totalCellMemory := header.PageSize * uint16(unsafe.Sizeof(entity.PageCell{}))
			pageCellsRaw := make([]byte, totalCellMemory)
			readLen, err = f.ReadAt(pageCellsRaw, int64(len(pageHeaderRaw)))
		}
		return &entity.Page{
			Header: header,
			Cells:  &pageCells,
		}, nil
	} else {
		return nil, err
	}
}

func WritePage(path string, page *entity.Page, serializer *PageSerializer) error {
	bytes := serializer.SerializePage(page)
	err := ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
