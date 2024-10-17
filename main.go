package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var config = LoadConfig()
var scripts = getScripts(config.getOrCreateScriptsDir())

func main() {
	if len(os.Args) == 1 {
		cmdrTui(scripts)
	} else {
		runScriptByName(os.Args[1])
	}
}

func cmdrTui(scripts []scriptFile) {
	for i, script := range scripts {
		fmt.Printf("[%d] %s\n", i, script.meta.Name())
	}

	fmt.Print("\n> ")

	var choice int = -1
	fmt.Scanln(&choice)
	if choice == -1 {
		fmt.Println("None selected")
		return
	}

	scripts[choice].run()
}

func runScriptByName(scriptName string) {
	maybeCommand := scriptName
	for _, script := range scripts {
		if maybeCommand == script.name || maybeCommand == script.meta.Name() {
			script.run()
			return
		}
	}
	fmt.Printf("'%s' is not a valid script name\n", maybeCommand)
}

type scriptFile struct {
	name string
	path string
	meta fs.FileInfo
	kind string
}

func (s scriptFile) run() {
	scriptPath := filepath.Join(s.path, s.meta.Name())
	args := []string{scriptPath}
	if len(os.Args) > 2 {
		args = append(args, os.Args[2:]...)
	}
	switch s.kind {
	case "py":
		execCommand("python3", []string{}, args)

	case "js":
		execCommand("node", []string{}, args)

	case "go":
		execCommand("go", []string{"run"}, args)

	case "sh":
		fallthrough
	default:
		execCommand("/bin/bash", []string{}, args)
	}
}

func execCommand(runtime string, runtimeArgs []string, cmdArgs []string) {
	args := slices.Concat(runtimeArgs, cmdArgs)
	cmd := exec.Command(runtime, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

func getScripts(path string) []scriptFile {
	entries, err := os.ReadDir(path)
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
			fileNameParts := strings.Split(entry.Name(), ".")
			extension := ""
			if len(fileNameParts) == 2 {
				extension = fileNameParts[1]
			}
			// skip hidden files
			if fileNameParts[0] == "" {
				continue
			}
			files = append(files, scriptFile{name: fileNameParts[0], path: path, meta: info, kind: extension})
		}
	}

	return files
}
