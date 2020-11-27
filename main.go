package main

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data"
	"ZeitDB/storage/repository"
	"encoding/json"
)

func main() {
	config := entity.Configuration{
		MetaInfoFilePath: "./dbj2",
	}
	dataSource := data.ConfigFileDataSource{}
	configManager := repository.ConfigRepository{}

	dataSource.SetConfig(&config)
	configManager.SetDataSource(dataSource)

	_, err := configManager.Initialize()

	configManager.UpdateMetaInfo(&entity.MetaInfo{
		GlobalLowestTimeStamp:  100,
		GlobalHighestTimeStamp: 10000,
		AmountOfPages:          30,
		AmountOfCells:          20,
		Version:                5,
	})

	metaInfo, err := configManager.ObtainMetaInfo()
	if err != nil {
		panic(err)
	}
	printMetaInfo(metaInfo)
}

func printMetaInfo(info *entity.MetaInfo) {
	jsonRepr, _ := json.Marshal(info)

	println(jsonRepr)
}
