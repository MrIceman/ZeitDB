package main

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data/page"
	"ZeitDB/storage/factory"
	"ZeitDB/storage/repository"
	"fmt"
)

func createEmptyPage(
	meta *entity.MetaInfo,
	configRepository *repository.ConfigRepository,
	pageRepository *repository.PageRepository,
) {

	result, err := pageRepository.CreateNewPage()

	if err != nil {
		panic(err)
	}

	meta.AmountOfPages += 1
	metaInfo, err := configRepository.UpdateMetaInfo(meta)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created page ", result.Header.PageNumber)
	fmt.Println("New page size: ", metaInfo.AmountOfPages)
}

func main() {
	configRepository := factory.CreateConfigRepository(
		"./dbj2",
		"./",
	)
	info, err := configRepository.Initialize()

	if err != nil {
		panic(err)
	}

	ds := page.PageFileDataSource{}
	ds.Init(info, configRepository.Config())
	pageRepository := repository.PageRepository{}
	pageRepository.SetDataSource(&ds)

	createEmptyPage(
		info,
		configRepository,
		&pageRepository,
	)

	for i := 0; i < int(info.AmountOfPages); i++ {
		createdPage, err := pageRepository.GetPage(int8(i))
		if err != nil {
			panic(err)
		}
		println(createdPage.Header.PageNumber)
	}
}
