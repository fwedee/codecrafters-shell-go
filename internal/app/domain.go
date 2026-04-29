package app

import (
	"fmt"
	"io"
)

type LineReader interface {
	ReadLine(prompt string) (string, error)
}

type Command func(args []string)

type shellInput struct {
	Name string
	Args []string
}

type App struct {
	in       LineReader
	out      io.Writer
	registry map[string]Command
}

func (a *App) ReadLine(prompt string) (string, error) {
	return a.in.ReadLine(prompt)
}

func (a *App) Print(msg string) {
	fmt.Fprint(a.out, msg)
}

func (a *App) Println(msg string) {
	fmt.Fprintln(a.out, msg)
}
