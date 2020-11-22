package main

import (
	"ZeitDB/storage/data"
	"ZeitDB/storage/manager"
	"ZeitDB/storage/model"
	"encoding/json"
)

func main() {
	config := model.Configuration{
		MetaInfoFilePath: "./db2",
	}
	configManager := manager.ConfigManager{
		DataSource: data.ConfigDataSource{
			Config: &config,
		}}

	_, err := configManager.Initialize()

	configManager.UpdateMetaInfo(&model.MetaInfo{
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

func printMetaInfo(info *model.MetaInfo) {
	jsonRepr, _ := json.Marshal(info)

	println(jsonRepr)
}
