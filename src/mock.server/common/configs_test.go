package common

import (
	"testing"
)

func TestInitConfigs(t *testing.T) {
	if err := InitConfigs(); err != nil {
		t.Fatal(err)
	}
	t.Logf("configs: %+v", RunConfigs)
}
