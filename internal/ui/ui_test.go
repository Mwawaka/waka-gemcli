package ui

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_PrintError(t *testing.T) {
	err := fmt.Errorf("connection failed")
	buf := &bytes.Buffer{}
	expectedText := "[ERROR] an error occurred: <connection failed>"
	terminalUi := NewUI(os.Stdin, os.Stdout, buf)
	terminalUi.PrintError(err)

	if !strings.Contains(buf.String(), expectedText) {
		t.Errorf("expected output to contain %q, got %q", expectedText, buf.String())
	}
}

func Test_PrintChunk(t *testing.T) {
	buf := &bytes.Buffer{}
	terminalUi := NewUI(os.Stdin, buf, os.Stderr)
	chunk := "This is the correct response"
	terminalUi.PrintChunk(chunk)
	if !strings.Contains(buf.String(), chunk) {
		t.Errorf("expected output to contain %q, got %q", chunk, buf.String())
	}
}
