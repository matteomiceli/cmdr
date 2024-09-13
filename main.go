package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	scripts := getScripts()
	for i, script := range scripts {
		fmt.Printf("[%d] %s\n", i, script.meta.Name())
	}

	fmt.Print("\nSelect a script to run: ")
	var choice int
	fmt.Scanln(&choice)

	fmt.Printf("\033[2J")

	scripts[choice].run()
}

type scriptFile struct {
	meta fs.FileInfo
	kind string
}

func (s scriptFile) run() {
	scriptPath := filepath.Join(getScriptsDir(), s.meta.Name())
	switch s.kind {
	case "py":
		runCommand("python3", scriptPath)

	case "js":
		runCommand("node", scriptPath)

	case "sh":
		fallthrough
	default:
		runCommand("/bin/bash", scriptPath)
	}
}

func runCommand(runtime string, args ...string) {
	cmd := exec.Command(runtime, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal()
	}
	cmd.Wait()
}

func getScripts() []scriptFile {
	entries, err := os.ReadDir(getScriptsDir())
	if err != nil {
		log.Fatal()
	}

	files := []scriptFile{}
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				log.Fatal()
			}
			scriptType := strings.Split(entry.Name(), ".")[1]
			files = append(files, scriptFile{meta: info, kind: scriptType})
		}
	}

	return files
}

func getScriptsDir() string {
	HOME_DIR, err := os.UserHomeDir()
	if err != nil {
		log.Fatal()
	}

	SCRIPTS_DIR := filepath.Join(HOME_DIR, "Documents", "scripts")
	if err != nil {
		log.Fatal()
	}

	return SCRIPTS_DIR
}
