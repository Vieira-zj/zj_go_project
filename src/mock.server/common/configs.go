package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const configFile = "mock_conf.json"

// Configs mock server configs.
type Configs struct {
	Meta   string        `json:"meta"`
	RunEnv string        `json:"run_env"`
	Server ServerConfigs `json:"server"`
}

// ServerConfigs server configs.
type ServerConfigs struct {
	RedisURI string `json:"redis_uri"`
}

// RunConfigs stores configs of mock server.
var RunConfigs Configs = Configs{
	RunEnv: "test", // test, prod
	Server: ServerConfigs{
		RedisURI: "http://localhost:6379",
	},
}

// InitConfigs reads mock server configs from cur directory.
func InitConfigs() error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	log.Println("load configs:", data)
	return json.Unmarshal(data, &RunConfigs)
}
