package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	_, err := GetConfig("../resource/config.json")
	if err != nil {
		t.Error("Initialize Config Failed\n")
	}
}
