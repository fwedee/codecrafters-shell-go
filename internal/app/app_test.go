package app_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/fwedee/codecrafters-shell-go/internal/app"
	"github.com/fwedee/codecrafters-shell-go/internal/cli"
)

func TestRunEcho(t *testing.T) {
	var out bytes.Buffer
	input := strings.NewReader("echo hello world\n")

	shell := app.NewApp(cli.NewInput(input, &out), &out)
	shell.Run()

	want := "$ hello world \n$ "
	if got := out.String(); got != want {
		t.Fatalf("output mismatch:\nwant %q\n got %q", want, got)
	}
}
