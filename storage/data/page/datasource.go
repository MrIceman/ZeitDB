package page

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data/page/io"
	"strconv"
)

type PageFileDataSource struct {
	config   *entity.Configuration
	metaInfo *entity.MetaInfo
}

func (p *PageFileDataSource) Init(metaInfo *entity.MetaInfo, configuration *entity.Configuration) {
	p.metaInfo = metaInfo
	p.config = configuration
}

func (p *PageFileDataSource) GetLastPage() *entity.Page {
	panic("implement me")
}

func (p *PageFileDataSource) GetFirstPage() *entity.Page {
	panic("implement me")
}

func (p *PageFileDataSource) CreateNewPage() (*entity.Page, error) {
	pageIndex := p.metaInfo.AmountOfPages + 1
	// get the last page
	strAmountOfPages := strconv.Itoa(int(pageIndex))
	filePath := p.config.PageRootFilePath + " " + strAmountOfPages

	// todo keyIndex has to be calculated differently, it needs to be the first
	// available page
	page := CreateEmptyPage(int8(pageIndex), int8(pageIndex))

	return page, io.WritePage(filePath,
		page,
		&io.PageSerializer{})
}
