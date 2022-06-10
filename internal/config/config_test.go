package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("./configss.json")
	fmt.Println(Config)
}

func TestExportConfig(t *testing.T) {
	//ExportDefaultConfig("./configs.json")
}
