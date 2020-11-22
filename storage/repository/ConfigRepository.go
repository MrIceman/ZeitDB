package repository

import (
	"ZeitDB/entity/model"
	"ZeitDB/storage/data"
)

type ConfigRepository struct {
	dataSource data.ConfigFileDataSource
}

func (c *ConfigRepository) SetDataSource(source data.ConfigFileDataSource) {
	c.dataSource = source
}

func (c *ConfigRepository) ObtainMetaInfo() (*model.MetaInfo, error) {
	return c.dataSource.GetMetaInfo()
}

func (c *ConfigRepository) Initialize() (*model.MetaInfo, error) {
	return c.dataSource.Init()
}

func (c *ConfigRepository) UpdateMetaInfo(metaInfo *model.MetaInfo) (*model.MetaInfo, error) {
	err := c.dataSource.SaveMetaInfo(metaInfo)
	if err != nil {
		return nil, err
	}
	return metaInfo, nil

}
