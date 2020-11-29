package repository

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data/page"
)

type PageRepository struct {
	dataSource *page.PageFileDataSource
}

func (pr *PageRepository) SetDataSource(dataSource *page.PageFileDataSource) {
	pr.dataSource = dataSource
}

func (pr *PageRepository) CreateNewPage() (*entity.Page, error) {
	page, err := pr.dataSource.CreateNewPage()

	return page, err
}

func (pr *PageRepository) GetPageHeader(pageNumber int) (*entity.PageHeader, error) {
	panic("Not implemented")
}
