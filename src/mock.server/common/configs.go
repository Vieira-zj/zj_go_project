package common

import (
	"encoding/json"
	"io/ioutil"
)

// ServerConfigs struct for server configs.
type ServerConfigs struct {
	RedisURI string `json:"redis_uri"`
}

// Configs configs for mock server.
type Configs struct {
	Meta   string        `json:"meta"`
	Server ServerConfigs `json:"server"`
}

// RunConfigs stores configs for mock server.
var RunConfigs Configs

// InitConfigs reads mock server configs from cur dir.
func InitConfigs() error {
	configsFile := "mock_conf.json"
	data, err := ioutil.ReadFile(configsFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &RunConfigs); err != nil {
		return err
	}
	return nil
}
