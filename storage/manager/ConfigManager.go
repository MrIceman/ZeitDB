package manager

import (
	"ZeitDB/storage/data"
	"ZeitDB/storage/model"
)

type ConfigManager struct {
	DataSource data.ConfigDataSource
}

func (c *ConfigManager) ObtainMetaInfo() (*model.MetaInfo, error) {
	return c.DataSource.GetMetaInfo()
}

func (c *ConfigManager) Initialize() (*model.MetaInfo, error) {
	return c.DataSource.Init()
}

func (c *ConfigManager) UpdateMetaInfo(metaInfo *model.MetaInfo) (*model.MetaInfo, error) {
	err := c.DataSource.SaveMetaInfo(metaInfo)
	if err != nil {
		return nil, err
	}
	return metaInfo, nil

}
