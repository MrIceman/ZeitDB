package page

import (
	"ZeitDB/entity"
	"strconv"
)

func CreateEmptyPage(pageNumber int8,
	keyIndex int8) *entity.Page {
	page := entity.Page{
		Header: &entity.PageHeader{
			PageNumber:       pageNumber,
			KeyIndex:         keyIndex,
			HighestTimeStamp: -1,
			LowestTimeStamp:  -1,
			PageSize:         0,
			IndexFileName:    "index_" + strconv.Itoa(int(pageNumber)),
		},
		Cells: &[]entity.PageCell{},
	}

	return &page

}
