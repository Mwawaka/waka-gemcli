package ui

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_PrintError(t *testing.T) {
	err := fmt.Errorf("connection failed")
	buf := &bytes.Buffer{}
	expectedText := "[ERROR] an error occurred: <connection failed>"
	terminalUi := NewUI(os.Stdout, buf)
	terminalUi.PrintError(err)
	
	if !strings.Contains(buf.String(), expectedText) {
		t.Errorf("expected output to contain %q, got %q", expectedText, buf.String())
	}
}

func Test_PrintResponse(t *testing.T) {
	buf := &bytes.Buffer{}
	terminalUi := NewUI(buf, os.Stderr)
	response := "This is the correct response"
	expectedText := fmt.Sprintf("\n🧠 GemCLI says:\n%s\n", "This is the correct response")
	terminalUi.PrintResponse(response, time.Microsecond)
	if !strings.Contains(buf.String(), expectedText) {
		t.Errorf("expected output to contain %q, got %q", expectedText, buf.String())
	}
}
