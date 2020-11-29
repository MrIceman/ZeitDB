package main

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data"
	"ZeitDB/storage/data/page"
	"ZeitDB/storage/repository"
	"fmt"
)

func createEmptyPage(config *entity.Configuration, meta *entity.MetaInfo, configRepository *repository.ConfigRepository) {
	ds := page.PageFileDataSource{}
	ds.Init(meta, config)
	pageRepository := repository.PageRepository{}
	pageRepository.SetDataSource(&ds)

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

	config := entity.Configuration{
		MetaInfoFilePath: "./dbj2",
	}
	dataSource := data.ConfigFileDataSource{}
	configRepository := repository.ConfigRepository{}

	dataSource.SetConfig(&config)
	configRepository.SetDataSource(dataSource)

	info, err := configRepository.Initialize()

	if err != nil {
		panic(err)
	}

	createEmptyPage(
		&config,
		info,
		&configRepository,
	)
}
