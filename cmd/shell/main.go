package main

import (
	"os"

	"github.com/fwedee/codecrafters-shell-go/internal/app"
	"github.com/fwedee/codecrafters-shell-go/internal/cli"
)

func main() {
	shell := app.NewApp(cli.NewInput(os.Stdin, os.Stdout), os.Stdout)
	shell.Run()
}
