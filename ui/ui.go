package ui

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

type UI struct {
	out io.Writer
	err io.Writer
}

func NewUI() *UI {
	return &UI{
		out: os.Stdout,
		err: os.Stderr,
	}
}

// PrintPrompt prints the styled  ~> prompt
func (u *UI) PrintPrompt() {
	neonGreen := color.RGB(57, 255, 20).FprintFunc()
	neonGreen(u.out, "⚡GemCLI ~> ")

}

// PrintResponse prints the response and adds typing animation
func (u *UI) PrintResponse(response string, delay time.Duration) {
	response = fmt.Sprintf("\n🧠 GemCLI says:\n%s\n", response)
	success := color.New(color.FgCyan).FprintFunc()

	for _, c := range response {
		success(u.out, string(c))
		time.Sleep(delay)
	}

}

// PrintError prints errors in a stylized format
func (u *UI) PrintError(err error) {
	failure := color.New(color.FgRed).FprintfFunc()

	failure(u.err, "[ERROR] an error occurred: <%v>\n", err)
}

// Start showcases an ascii logo with a welcome message
func (u *UI) Start() {
	logo := `
 ██████╗ ███████╗███╗   ███╗ ██████╗██╗     ██╗
██╔════╝ ██╔════╝████╗ ████║██╔════╝██║     ██║
██║  ███╗█████╗  ██╔████╔██║██║     ██║     ██║
██║   ██║██╔══╝  ██║╚██╔╝██║██║     ██║     ██║
╚██████╔╝███████╗██║ ╚═╝ ██║╚██████╗███████╗██║
 ╚═════╝ ╚══════╝╚═╝     ╚═╝ ╚═════╝╚══════╝╚═╝ ...`

	magenta := color.New(color.FgMagenta).FprintlnFunc()
	magenta(u.out, "Welcome to ")
	magenta(u.out, logo)
}
