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
	"strings"
)

//go:embed defaultConfig.json
var defaultConfig []byte

// builtins

//go:embed builtins/_new.sh
var _new []byte

//go:embed builtins/_edit.sh
var _edit []byte

//go:embed builtins/_rm.sh
var _rm []byte

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

	_, err := os.Stat(scriptsDir)
	// if scripts dir already exists
	if err == nil {
		c.ScriptsPath = scriptsDir
		return scriptsDir
	}

	fmt.Printf("Scripts directory not found, creating one at %s \n\n", scriptsDir)
	c.ScriptsPath = scriptsDir
	os.MkdirAll(scriptsDir, os.ModePerm)

	fmt.Println("Would you like to add builtin scripts? (y)es / (n)o?")
	fmt.Println("For the best out of the box experience, it's recommended that you also install builtins.")

	var ans string
	fmt.Scan(&ans)
	ans = strings.TrimSpace(ans)
	if ans == "y" || ans == "yes" {
		os.WriteFile(path.Join(scriptsDir, "_new.sh"), _new, 0644)
		os.WriteFile(path.Join(scriptsDir, "_edit.sh"), _edit, 0644)
		os.WriteFile(path.Join(scriptsDir, "_rm.sh"), _rm, 0644)
	}

	return scriptsDir
}

func LoadConfig() Config {
	var config Config
	err := json.Unmarshal(getOrCreateConfigFile(), &config)
	if err != nil {
		log.Fatal(err)
	}

	config.getOrCreateScriptsDir()

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

		fmt.Print(paint("green", "Config file not found, generated default in "))
		fmt.Printf(paint("magenta", "%s\n\n"), cmdrConfigPath)
	}

	config, err := os.ReadFile(cmdrConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
