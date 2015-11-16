package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"runtime"
)

var (
	errNilConfig = errors.New("Config object is empty.")
)

func NewConfig() error {
	configurationFile := "config.json"
	err := loadJsonFile(configurationFile, &Config)
	if err != nil {
		return err
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	return err
}

func loadJsonFile(filename string, config interface{}) error {
	if config == nil {
		return errNilConfig
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, config)
}
