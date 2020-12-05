package io

import "ZeitDB/entity"

type IndexSerializer struct {
}

func (is *IndexSerializer) SerializeIndexPage(page entity.IndexPage) []byte {
	return []byte{}
}

func (is *IndexSerializer) serializeIndexEntry(entry entity.IndexEntry) *[]byte {
	return &[]byte{}
}
