package cli

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct{}

func (Input) Get(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

type Output struct{}

func (Output) Print(text string) {
	fmt.Print(text)
}

func (Output) Println(text string) {
	fmt.Println(text)
}
