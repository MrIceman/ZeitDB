package io

import (
	"ZeitDB/entity"
	"io/ioutil"
	"strconv"
)

func ReadPage(path string) *entity.Page {
	return nil
}

func WritePage(path string, page *entity.Page, serializer *PageSerializer) error {
	bytes := serializer.SerializePage(page)
	fullPath := path + strconv.Itoa(int(page.Header.PageNumber))
	err := ioutil.WriteFile(fullPath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
