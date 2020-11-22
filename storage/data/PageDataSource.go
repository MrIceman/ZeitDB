package data

import "ZeitDB/storage/model"

type PageSource interface {
	SetConfiguration(info *model.MetaInfo)
	GetLastPage() *model.Page
	GetFirstPage() *model.Page
	CreateNewPage() *model.Page
}

type PageDataSource struct {
	metaInfo *model.MetaInfo
}

func (p PageDataSource) SetConfiguration(info *model.MetaInfo) {
	p.metaInfo = info
}

func (p PageDataSource) GetLastPage() *model.Page {
	panic("implement me")
}

func (p PageDataSource) GetFirstPage() *model.Page {
	panic("implement me")
}

func (p PageDataSource) CreateNewPage() *model.Page {
	panic("implement me")
}
