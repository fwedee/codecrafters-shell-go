package app

import (
	"os/exec"
)

func findExecutable(commandName string) (string, bool) {
	path, err := exec.LookPath(commandName)
	if err != nil {
		return "", false
	}

	return path, true
}

func (a *App) executeExternal(input shellInput) {
	if _, exists := findExecutable(input.Name); !exists {
		a.Println(input.Name + ": command not found")
		return
	}

	cmd := exec.Command(input.Name, input.Args...)
	cmd.Stdout = a.out

	if err := cmd.Run(); err != nil {
		a.Println(err.Error())
		return
	}
}
