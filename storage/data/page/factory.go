package page

import "ZeitDB/entity/model"

func CreateEmptyPage(pageNumber uint16, keyIndex int8) *model.Page {
	page := model.Page{
		Header: model.PageHeader{
			PageNumber:       pageNumber,
			KeyIndex:         keyIndex,
			HighestTimeStamp: 0,
			LowestTimeStamp:  0,
		},
		Cells: []model.PageCell{},
	}

	return &page

}
