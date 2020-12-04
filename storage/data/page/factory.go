package page

import (
	"ZeitDB/entity"
)

func CreateEmptyPage(pageNumber int8, keyIndex int8) *entity.Page {
	page := entity.Page{
		Header: &entity.PageHeader{
			PageNumber:       pageNumber,
			KeyIndex:         keyIndex,
			HighestTimeStamp: -1,
			LowestTimeStamp:  -1,
		},
		Cells: &[]entity.PageCell{},
	}

	return &page

}
