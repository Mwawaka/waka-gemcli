![GemCLI](assets/gemcli.png)
# ✨ Simple Gemini CLI Client

> A minimal, clean, and hackable Command Line Interface for interacting with Gemini on your terminal.

---

## 🖥️ Demo

```bash
⚡ GemCLI ~> "Tell me a programming joke"

🧠 GemCLI says:
Why do programmers like dark mode? Because light attracts bugs 😂😂.
```

---

```go
// cmd/root.go
func Execute() {
    rootCmd.Execute()
}
Then main.go simply calls:
go// main.go
func main() {
    cmd.Execute()
}
```