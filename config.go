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
	"path/filepath"
	"slices"
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

func (c *Config) getRunner(fileExtension string) (Runner, error) {
	for _, runner := range c.Runners {
		if slices.Contains(runner.FileAssociations, fileExtension) {
			return runner, nil
		}
	}
	return Runner{}, errors.New("No runner associated with this file.")
}

func (c *Config) getOrCreateScriptsDir() string {
	scriptsDir := ""
	if c.ScriptsPath != "" {
		scriptsDir = c.ScriptsPath
	} else {
		// if not defined in config, use default location
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		scriptsDir = filepath.Join(home, "scripts")
	}

	os.MkdirAll(scriptsDir, os.ModePerm)

	return scriptsDir
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
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(paint("green", "Couldn't find config file, generated default config in "))
		fmt.Printf(paint("default", "%s\n\n"), cmdrConfigPath)
	}

	config, err := os.ReadFile(cmdrConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
