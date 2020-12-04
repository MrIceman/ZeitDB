package factory

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data"
	"ZeitDB/storage/repository"
)

func CreateConfigRepository(
	metaInfoFilepath string,
	pageRootFilePath string) *repository.ConfigRepository {

	config := entity.Configuration{
		MetaInfoFilePath: metaInfoFilepath,
		PageRootFilePath: pageRootFilePath,
	}
	dataSource := data.ConfigFileDataSource{}
	configRepository := repository.ConfigRepository{}

	dataSource.SetConfig(&config)
	configRepository.SetDataSource(dataSource)

	return &configRepository
}
