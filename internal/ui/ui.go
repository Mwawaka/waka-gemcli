package ui

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
)

type UI struct {
	out io.Writer
	err io.Writer
}

func NewUI(out, err io.Writer) *UI {
	return &UI{
		out: out,
		err: err,
	}
}

// PrintPrompt prints the styled  ~> prompt
func (u *UI) PrintPrompt() {
	neonGreen := color.RGB(57, 255, 20).FprintFunc()
	neonGreen(u.out, "\n⚡GemCLI ~> ")
}

// PrintResponse prints the response and adds typing animation
func (u *UI) PrintResponse(response string, delay time.Duration) {
	response = fmt.Sprintf("\n🧠 GemCLI says:\n%s\n\n", response)
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

// PrintHelp prints supported commands and their descriptions

func (u *UI) PrintHelp() {
	neonGreen := color.RGB(57, 255, 20).FprintlnFunc()
	w := tabwriter.NewWriter(u.out, 0, 0, 3, ' ', 0)
	neonGreen(w, "COMMAND\tDESCRIPTION\n")
	neonGreen(w, "!clear\t Resets the chat history")
	neonGreen(w, "!exit\t Terminates the cli")
	neonGreen(w, "!help\t Displays the command table")
	w.Flush() //Prints the formatted table
}

// PrintInvalidInput notifies user that a command is not supported
func (u *UI) PrintInvalidInput(command string) {
	yellow := color.New(color.FgYellow).FprintfFunc()
	yellow(u.out, "[UNKNOWN] %s is not a valid command. Type !help for available commands\n", command)
}
