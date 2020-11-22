package page

import (
	"ZeitDB/entity/model"
	"ZeitDB/storage/data/page/page_utils"
	"strconv"
)

type PageSource interface {
	SetConfiguration(info *model.Configuration)
	GetLastPage() *model.Page
	GetFirstPage() *model.Page
	CreateNewPage() *model.Page
	SearchCell()
}

type PageFileDataSource struct {
	config   *model.Configuration
	metaInfo *model.MetaInfo
}

func (p PageFileDataSource) Init(metaInfo *model.MetaInfo, configuration *model.Configuration) {
	p.metaInfo = metaInfo
	p.config = configuration
}

func (p PageFileDataSource) GetLastPage() *model.Page {
	panic("implement me")
}

func (p PageFileDataSource) GetFirstPage() *model.Page {
	panic("implement me")
}

func (p PageFileDataSource) CreateNewPage() *model.Page {
	amountOfPages := p.metaInfo.AmountOfPages
	// get the last page
	strAmountOfPages := strconv.Itoa(amountOfPages)
	filePath := p.config.PageRootFilePath + " " + strAmountOfPages

	return page_utils.ReadPage(filePath)
}
