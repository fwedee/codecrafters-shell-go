package app

import "io"

func NewApp(in LineReader, out io.Writer) *App {
	a := &App{
		in:       in,
		out:      out,
		registry: map[string]Command{},
	}

	a.registerBuiltins()
	return a
}

func (a *App) Run() {
	for {
		rawInput, err := a.ReadLine("$ ")
		if err != nil {
			return
		}

		input := parseInput(rawInput)

		if input.Name == "" {
			continue
		}

		if command, ok := a.registry[input.Name]; ok {
			command(input.Args)
			continue
		}

		a.executeExternal(input)
	}
}
