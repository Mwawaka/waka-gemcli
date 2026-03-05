## Phase 1
Built:
- ✅ Cyberpunk styled UI package with PrintPrompt, PrintResponse, PrintError, Start
- ✅ UI struct with dependency injected writers
- ✅ Tests for PrintError and PrintResponse

Go concepts covered:

- Struct methods vs functions
- Encapsulation with unexported fields
- Dependency injection via io.Writer
- Error wrapping with %w
- for range over strings and runes
- Writing tests with testing package
- Using bytes.Buffer to capture output in tests

---

## Phase 2

- ✅ Chat mode with conversation history via genai.Chat
- ✅ !exit command with labeled break loop
- ✅ !clear command with safe chat reset
- ✅ !help command with tabwriter formatted output
- ✅ Unknown command handling with PrintInvalidInput