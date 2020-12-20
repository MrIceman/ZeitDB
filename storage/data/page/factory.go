package page

import (
	"ZeitDB/entity"
)

func CreateEmptyPage(pageNumber int8,
	keyIndex int8) *entity.Page {

	page := entity.Page{
		Header: &entity.PageHeader{
			Magic:            5,
			PageNumber:       pageNumber,
			KeyIndex:         keyIndex,
			HighestTimeStamp: -1,
			LowestTimeStamp:  -1,
			PageSize:         0,
		},
		Cells: &[]entity.PageCell{},
	}

	return &page

}
