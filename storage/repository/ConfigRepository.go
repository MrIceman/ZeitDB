package repository

import (
	"ZeitDB/entity"
	"ZeitDB/storage/data"
)

type ConfigRepository struct {
	dataSource data.ConfigFileDataSource
}

func (c *ConfigRepository) SetDataSource(source data.ConfigFileDataSource) {
	c.dataSource = source
}

func (c *ConfigRepository) ObtainMetaInfo() (*entity.MetaInfo, error) {
	return c.dataSource.GetMetaInfo()
}

func (c *ConfigRepository) Initialize() (*entity.MetaInfo, error) {
	return c.dataSource.Init()
}

func (c *ConfigRepository) UpdateMetaInfo(metaInfo *entity.MetaInfo) (*entity.MetaInfo, error) {
	err := c.dataSource.SaveMetaInfo(metaInfo)
	if err != nil {
		panic(err)
		return nil, err
	}
	return metaInfo, nil

}
