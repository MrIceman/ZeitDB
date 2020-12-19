package main

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data/page"
	"ZeitDB/storage/factory"
	"ZeitDB/storage/repository"
	"fmt"
)

func createEmptyPage(
	metaInfo *entity.MetaInfo,
	configRepository *repository.ConfigRepository,
	pageRepository *repository.PageRepository,
) {

	result, err := pageRepository.CreateNewPage()

	if err != nil {
		panic(err)
	}

	metaInfo.AmountOfPages += 1
	updatedMetaInfo, err := configRepository.UpdateMetaInfo(metaInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created page ", result.Header.PageNumber)
	fmt.Println("New page size: ", updatedMetaInfo.AmountOfPages)
}

func main() {
	configRepository := factory.CreateConfigRepository(
		"./dbj2",
		"./",
	)
	metaInfo, err := configRepository.InitializeMetaInfo()

	if err != nil {
		panic(err)
	}

	ds := page.PageFileDataSource{}
	ds.Init(metaInfo, configRepository.Config())
	pageRepository := repository.PageRepository{}
	pageRepository.SetDataSource(&ds)
	createEmptyPage(
		metaInfo,
		configRepository,
		&pageRepository,
	)
	for i := 0; i < int(metaInfo.AmountOfPages); i++ {
		createdPage, err := pageRepository.GetPage(int8(i))
		if err != nil {
			panic(err)
		} else {
			println(createdPage.Header)
		}
	}
}
