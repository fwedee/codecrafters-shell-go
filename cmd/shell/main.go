package main

import (
	"github.com/fwedee/codecrafters-shell-go/internal/app"
	"github.com/fwedee/codecrafters-shell-go/internal/cli"
)

func main() {

	app := app.NewApp(cli.Input{}, cli.Output{})
	app.Run()
}
