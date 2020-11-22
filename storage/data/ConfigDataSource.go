package data

import (
	"ZeitDB/storage/model"
	"io/ioutil"
	"os"
)

type ConfigSource interface {
	SetConfig(path string) error
	GetMetaInfo() (*model.MetaInfo, error)
	SetMetaInfo(info *model.MetaInfo) error
	Init() (*model.MetaInfo, error)
}

type ConfigDataSource struct {
	Config *model.Configuration
}

func (c *ConfigDataSource) SetConfig(config *model.Configuration) error {
	c.Config = config
	return nil
}

func (c *ConfigDataSource) GetMetaInfo() (*model.MetaInfo, error) {
	f, err := os.Open(c.Config.MetaInfoFilePath)
	if err == nil {
		defer f.Close()
		byteArray := make([]byte, 20)
		_, err = f.Read(byteArray)
		if err == nil {
			return model.FromByteArray(byteArray), nil
		}
		return nil, err
	}
	println("shit err is not null")
	println("file:", f)
	return nil, err
}

func (c *ConfigDataSource) Init() (*model.MetaInfo, error) {
	res, _ := c.GetMetaInfo()
	if res != nil {
		println("ConfigManager already exists")
		return res, nil
	}
	println("Creating a new ConfigManager.")
	metaInfo := model.MetaInfo{Version: 001}
	err := c.SaveMetaInfo(&metaInfo)
	if err != nil {
		return nil, err
	}
	createdMetaInfo, err := c.GetMetaInfo()
	return createdMetaInfo, err
}

/**
Writes a new MetaInfo to disk
*/
func (c *ConfigDataSource) SaveMetaInfo(info *model.MetaInfo) error {
	path := c.Config.MetaInfoFilePath
	infoBytes := info.ToByteArray()
	err := ioutil.WriteFile(path, infoBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
