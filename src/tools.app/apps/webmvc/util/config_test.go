package util

import (
	"os"
	"path/filepath"
	"testing"
)

var fpath = filepath.Join(os.Getenv("GOPATH"), "src/tools.app/apps/webmvc/app.dev.conf")

func TestParseConfig(t *testing.T) {
	config := new(Configs)
	err := config.Parse(fpath)
	if err != nil {
		t.Fatal("failed with error:", err.Error())
	}

	configs := config.GetAllCfg()
	t.Log("log configs:", configs["log"])
	t.Log("database configs:", configs["database"])
	t.Log("server configs:", configs["server"])
}

func TestGetSection(t *testing.T) {
	config := new(Configs)
	err := config.Parse(fpath)
	if err != nil {
		t.Fatal("failed with error:", err.Error())
	}

	t.Log(config.GetSection("database"))
	t.Log(config.GetSection("server"))
}
