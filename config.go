package main

import (
	_ "embed"
	"errors"
	"io/fs"
	"log"
	"os"
	"path"
)

//go:embed defaultConfig.json
var defaultConfig []byte

func LoadConfig() {
	getOrCreateConfigFile()
	// return unmarshalled config
}

func getOrCreateConfigFile() string {
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

	return cmdrConfigPath
}
