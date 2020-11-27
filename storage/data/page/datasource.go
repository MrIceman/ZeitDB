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

func (p *PageFileDataSource) Init(metaInfo *model.MetaInfo, configuration *model.Configuration) {
	p.metaInfo = metaInfo
	p.config = configuration
}

func (p *PageFileDataSource) GetLastPage() *model.Page {
	panic("implement me")
}

func (p *PageFileDataSource) GetFirstPage() *model.Page {
	panic("implement me")
}

func (p *PageFileDataSource) CreateNewPage() *model.Page {
	pageIndex := p.metaInfo.AmountOfPages + 1
	// get the last page
	strAmountOfPages := strconv.Itoa(pageIndex)
	filePath := p.config.PageRootFilePath + " " + strAmountOfPages

	// todo keyIndex has to be calculated differently, it needs to be the first
	// available page
	page := CreateEmptyPage(uint16(pageIndex), int8(pageIndex))

	return page_utils.ReadPage(filePath)
}
