package page

import (
	"ZeitDB/entity/model"
	"ZeitDB/entity/model/page"
	"ZeitDB/storage/data/page/page_utils"
	"strconv"
)

type PageSource interface {
	SetConfiguration(info *model.Configuration)
	GetLastPage() *page.Page
	GetFirstPage() *page.Page
	CreateNewPage() *page.Page
	SearchCell()
}

type PageFileDataSource struct {
	config   *model.Configuration
	metaInfo *model.MetaInfo
}

func (p *PageFileDataSource) Init(metaInfo *model.MetaInfo, configuration *model.Configuration) {
	p.metaInfo = metaInfo
	p.config = configuration
}

func (p *PageFileDataSource) GetLastPage() *page.Page {
	panic("implement me")
}

func (p *PageFileDataSource) GetFirstPage() *page.Page {
	panic("implement me")
}

func (p *PageFileDataSource) CreateNewPage() *page.Page {
	pageIndex := p.metaInfo.AmountOfPages + 1
	// get the last page
	strAmountOfPages := strconv.Itoa(pageIndex)
	filePath := p.config.PageRootFilePath + " " + strAmountOfPages

	// todo keyIndex has to be calculated differently, it needs to be the first
	// available page
	page := CreateEmptyPage(uint16(pageIndex), int8(pageIndex))

	return page_utils.ReadPage(filePath)
}
