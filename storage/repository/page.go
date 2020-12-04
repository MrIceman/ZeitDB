package repository

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data/page"
)

/**
PageRepository is used to fetch data from the [page.PageFileDataSource]
or from a local cache, to avoid multiple file i/o's.
*/
type PageRepository struct {
	dataSource *page.PageFileDataSource
}

func (pr *PageRepository) SetDataSource(dataSource *page.PageFileDataSource) {
	pr.dataSource = dataSource
}

func (pr *PageRepository) CreateNewPage() (*entity.Page, error) {
	result, err := pr.dataSource.CreateNewPage()

	return result, err
}

func (pr *PageRepository) GetPageHeader(pageNumber int) (*entity.PageHeader, error) {
	panic("Not implemented")
}

func (pr *PageRepository) GetPage(pageNumber int8) (*entity.Page, error) {
	return pr.dataSource.GetPage(pageNumber)
}
