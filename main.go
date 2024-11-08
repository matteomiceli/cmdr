package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
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
	fmt.Println("Scripts:\n--------")
	numSelect := 0
	for _, script := range scripts {
		if !script.hidden {
			fmt.Printf("[%d] %s\n", numSelect, script.meta.Name())
			numSelect++
		}
	}

	fmt.Print(paint("cyan", "\n> "))
	args := captureInput()

	if len(args) == 0 {
		fmt.Println("None selected")
		return
	}

	choice, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Not a valid selection")
	}

	// user selection includes args ie.
	// > 0 --label test
	if len(args) > 1 {
		scripts[choice].run(args[1:]...)
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
	name   string
	path   string
	meta   fs.FileInfo
	kind   string
	hidden bool
}

func (s scriptFile) run(passedArgs ...string) {
	scriptPath := filepath.Join(s.path, s.meta.Name())
	args := []string{scriptPath}
	// Called from cli
	if len(os.Args) > 2 {
		args = append(args, os.Args[2:]...)
		// Called from TUI
	} else if len(passedArgs) > 0 {
		args = append(args, passedArgs...)
	}

	r, err := config.getRunner(s.kind)
	if err != nil {
		log.Fatal(err)
	}
	execCommand(r.Runtime, r.RuntimeArgs, args)
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
			var extension string
			if len(fileNameParts) == 2 {
				extension = fileNameParts[1]
			}
			// skip files hidden from file system (eg. .config)
			if fileNameParts[0] == "" {
				continue
			}
			isHidden := false
			formattedName := fileNameParts[0]
			if string(fileNameParts[0][0]) == "_" {
				isHidden = true
				// remove underscore so script can be targetted as argument
				// ie. _new can be called using `cmdr new`
				formattedName = fileNameParts[0][1:]
			}
			files = append(
				files, scriptFile{
					name:   formattedName,
					path:   path,
					meta:   info,
					kind:   extension,
					hidden: isHidden,
				},
			)
		}
	}

	return files
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

func captureInput() []string {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() && scanner.Text() != "" {
		input = strings.Split(scanner.Text(), " ")
	}
	return input
}
