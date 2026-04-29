package cli

import (
	"bufio"
	"fmt"
	"io"
)

type Input struct {
	in  *bufio.Scanner
	out io.Writer
}

func NewInput(in io.Reader, out io.Writer) *Input {
	return &Input{
		in:  bufio.NewScanner(in),
		out: out,
	}
}

func (i *Input) ReadLine(prompt string) (string, error) {
	if _, err := fmt.Fprint(i.out, prompt); err != nil {
		return "", err
	}

	if !i.in.Scan() {
		if err := i.in.Err(); err != nil {
			return "", err
		}

		return "", io.EOF
	}

	return i.in.Text(), nil
}
