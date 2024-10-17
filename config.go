package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
)

//go:embed defaultConfig.json
var defaultConfig []byte

type Runner struct {
	Runtime          string   `json:"runtime"`
	RuntimeArgs      []string `json:"runtimeArgs"`
	FileAssociations []string `json:"fileAssociations"`
}

type Config struct {
	Runners     []Runner `json:"runners"`
	ScriptsPath string   `json:"scriptsPath"`
}

func LoadConfig() Config {
	var config Config
	err := json.Unmarshal(getOrCreateConfigFile(), &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func getOrCreateConfigFile() []byte {
	osConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	cmdrConfigDir := path.Join(osConfigDir, "cmdr")
	err = os.MkdirAll(cmdrConfigDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	cmdrConfigPath := path.Join(cmdrConfigDir, "config.json")
	_, nofile := os.Stat(cmdrConfigPath)
	// config does not exist, so we make one
	if errors.Is(nofile, fs.ErrNotExist) {
		err = os.WriteFile(cmdrConfigPath, defaultConfig, 0644)
		if err == nil {
			log.Fatal(err)
		}
	}

	config, err := os.ReadFile(cmdrConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
