package app

type Input interface {
	Get(prompt string) string
}

type Output interface {
	Print(text string)
	Println(text string)
}

type Command func(args []string)

type App struct {
	in       Input
	out      Output
	registry map[string]Command
}

func (a *App) GetInput(prompt string) string {
	return a.in.Get(prompt)
}

func (a *App) Print(msg string) {
	a.out.Print(msg)
}

func (a *App) Println(msg string) {
	a.out.Println(msg)
}
