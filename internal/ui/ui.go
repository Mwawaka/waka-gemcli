package ui

import (
	"bufio"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/Mwawaka/waka-gemcli/internal/config"
	"github.com/fatih/color"
	"google.golang.org/genai"
)

type UI struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

func NewUI(in io.Reader, out, err io.Writer) *UI {
	return &UI{
		in:  in,
		out: out,
		err: err,
	}
}

// PrintPrompt prints the styled  ~> prompt
func (u *UI) PrintPrompt() {
	neonGreen := color.RGB(57, 255, 20).FprintFunc()
	neonGreen(u.out, "\n⚡GemCLI ~> ")
}

// PrintChunk prints streamed tokens
func (u *UI) PrintChunk(chunk string) {
	response := color.New(color.FgCyan).FprintFunc()
	response(u.out, chunk)
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

// PrintModelInfo prints details about all available models
func (u *UI) PrintModelInfo(models []*genai.Model) {
	var desc string
	neonGreen := color.RGB(57, 255, 20).FprintfFunc()
	w := tabwriter.NewWriter(u.out, 0, 0, 4, ' ', 0)
	neonGreen(w, "Name\tVersion\tThinking\tDescription\n")
	neonGreen(w, "----\t-------\t--------\t-----------\n")

	for _, model := range models {
		desc = model.Description
		if len(desc) > 100 {
			desc = desc[:100] + "..."
		}
		neonGreen(w, "%s\t%s\t%t\t%s\n", model.Name, model.Version, model.Thinking, desc)
	}

	w.Flush()
}

// PromptInitialConfig guides the user through initial configuration by collecting their API key and model preference
func (u *UI) PromptInitialConfig() (*config.Config, error) {
	reader := bufio.NewReader(u.in)
	magenta := color.New(color.FgMagenta).FprintlnFunc()
	u.Start()
	magenta(u.out, "\nNo config found. Let's get you set up!")
	magenta(u.out, "Enter your API key:")
	firstInput, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	magenta(u.out, "Enter your preferred model (default: gemini-2.5-flash):")
	secondInput, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	apiKey := strings.TrimSpace(firstInput)
	model := strings.TrimSpace(secondInput)

	if model == "" {
		model = "gemini-2.5-flash"
	}

	return &config.Config{
		APIKey: apiKey,
		Model:  model,
	}, nil

}
