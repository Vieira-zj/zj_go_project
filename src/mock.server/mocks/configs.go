package mocks

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ServerConfigs :
type ServerConfigs struct {
	Redis string `json:"redis"`
}

// Configs :
type Configs struct {
	Desc   string        `json:"desc"`
	Server ServerConfigs `json:"server"`
}

// RunConfigs : init from main.go
var RunConfigs Configs

// InitConfigs : read configs in cur dir
func InitConfigs() {
	configsFile := "mock_conf.json"
	data, err := ioutil.ReadFile(configsFile)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(data, &RunConfigs)
	if err != nil {
		log.Fatalln(err)
	}
}
