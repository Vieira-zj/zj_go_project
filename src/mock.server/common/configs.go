package common

import (
	"encoding/json"
	"io/ioutil"
)

// Configs mock server configs.
type Configs struct {
	Meta   string        `json:"meta"`
	Server ServerConfigs `json:"server"`
}

// ServerConfigs server configs.
type ServerConfigs struct {
	RedisURI string `json:"redis_uri"`
}

// RunConfigs stores configs of mock server.
var RunConfigs Configs

// InitConfigs reads mock server configs from cur directory.
func InitConfigs() error {
	data, err := ioutil.ReadFile("mock_conf.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &RunConfigs); err != nil {
		return err
	}
	return nil
}
