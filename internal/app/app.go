package app

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type ShellInput struct {
	Name string
	Args []string
}

var builtinCommands = map[string]struct{}{
	"exit": {},
	"echo": {},
	"type": {},
}

func NewApp(in Input, out Output) *App {
	a := &App{
		in:       in,
		out:      out,
		registry: map[string]Command{},
	}

	a.registry["exit"] = func(args []string) {
		os.Exit(0)
	}
	a.registry["echo"] = func(args []string) {
		for _, arg := range args {
			a.Print(arg + " ")
		}
		a.Println("")
	}
	a.registry["type"] = func(args []string) {
		if len(args) < 1 {
			fmt.Println("An argument for type is required.")
			return
		}
		_, exit := builtinCommands[args[0]]
		if exit {
			fmt.Println(args[0] + " is a shell builtin")
		} else {
			if exists, path := checkIfBinaryInPath(args[0]); exists {
				fmt.Println(args[0] + " is " + path + string(os.PathSeparator) + args[0])
			} else {
				fmt.Println(args[0] + ": not found")
			}
		}
	}
	return a
}

func checkIfBinaryInPath(commandName string) (bool, string) {
	rawPath := os.Getenv("PATH")
	paths := strings.SplitSeq(rawPath, string(os.PathListSeparator))

	for path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			//log.Fatal(err)
			continue
		}

		for _, file := range files {
			if commandName == file.Name() {
				info, err := os.Stat(path + string(os.PathSeparator) + file.Name())
				if err != nil {
					//log.Fatal(err)
					continue
				}
				mode := info.Mode()
				if mode&0111 != 0 {
					return true, path
				}
			}
		}
	}
	return false, ""
}

func (a *App) Run() {
	for {
		var input ShellInput
		rawInput := a.GetInput("$ ")

		fields := strings.Fields(rawInput)
		if len(fields) > 0 {
			input.Name = fields[0]
			input.Args = fields[1:]
		}

		if command, ok := a.registry[input.Name]; ok {
			command(input.Args)

		} else {
			if exists, _ := checkIfBinaryInPath(input.Name); exists {
				cmd := exec.Command(input.Name, input.Args...)

				var out bytes.Buffer
				cmd.Stdout = &out

				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}
				a.Print(out.String())
			} else {
				a.Println(input.Name + ": command not found")
			}
		}
	}
}

// Notes
// Run() = shell loop
// parser = converts line into command + args
// dispatcher/controller = finds matching command
// command = executes behavior
