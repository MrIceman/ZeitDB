package data

import (
	"ZeitDB/entity"
	"io/ioutil"
	"os"
)

type (
	ConfigSource interface {
		SetConfig(path string) error
		GetMetaInfo() (*entity.MetaInfo, error)
		SetMetaInfo(info *entity.MetaInfo) error
		Init() (*entity.MetaInfo, error)
	}
)

/**
	TODO extract all the file accessing into a
	separate class and generalize / abstract it
    to make it reusable
*/
type ConfigFileDataSource struct {
	config *entity.Configuration
}

func (c *ConfigFileDataSource) SetConfig(config *entity.Configuration) error {
	c.config = config
	return nil
}

func (c *ConfigFileDataSource) GetMetaInfo() (*entity.MetaInfo, error) {
	f, err := os.Open(c.config.MetaInfoFilePath)
	if err == nil {
		defer f.Close()
		byteArray := make([]byte, 20)
		_, err = f.Read(byteArray)
		if err == nil {
			return entity.FromByteArray(byteArray), nil
		}
		return nil, err
	}
	println("shit err is not null")
	println("file:", f)
	return nil, err
}

func (c *ConfigFileDataSource) Init() (*entity.MetaInfo, error) {
	res, _ := c.GetMetaInfo()
	if res != nil {
		println("ConfigManager already exists")
		return res, nil
	}
	println("Creating a new ConfigManager.")
	metaInfo := entity.MetaInfo{Version: 001}
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
func (c *ConfigFileDataSource) SaveMetaInfo(info *entity.MetaInfo) error {
	path := c.config.MetaInfoFilePath
	infoBytes := info.ToByteArray()
	err := ioutil.WriteFile(path, infoBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
