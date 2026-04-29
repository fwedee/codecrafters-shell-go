package app

import (
	"os"
	"path/filepath"
)

var builtinCommands = map[string]struct{}{
	"exit": {},
	"echo": {},
	"type": {},
}

func (a *App) registerBuiltins() {
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
			a.Println("An argument for type is required.")
			return
		}

		name := args[0]
		if _, exists := builtinCommands[name]; exists {
			a.Println(name + " is a shell builtin")
			return
		}

		path, exists := findExecutable(name)
		if exists {
			a.Println(name + " is " + filepath.Clean(path))
			return
		}

		a.Println(name + ": not found")
	}
}
