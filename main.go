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
	//RUNNER_CONTEXT, err := os.Getwd()
	//if err != nil {
	//	log.Fatal()
	//}

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
		fmt.Print(runCommand("python3", scriptPath))

	case "js":
		fmt.Print(runCommand("node", scriptPath))

	default:
		fmt.Println("Nothing")
	}
}

func runCommand(runtime string, args ...string) string {
	cmd := exec.Command(runtime, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
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
