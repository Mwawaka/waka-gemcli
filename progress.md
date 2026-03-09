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

## Phase 3 complete:

- ✅ Replaced SendMessage with SendMessageStream
- ✅ Callback pattern for chunk processing
- ✅ PrintChunk replaces PrintResponse
- ✅ Removed artificial typing animation — streaming provides natural flow
- ✅ Tests updated

### Go concepts covered in Phase 3:

- Iterators iter.Seq2 introduced in Go 1.23
- Callback functions as parameters
- Streaming vs buffered responses
- Removing code that no longer serves a purpose

## Phase 4 — Cobra rebuild complete:

- ✅ Project restructured with Cobra
- ✅ rootCmd with PersistentPreRunE for shared client/config initialisation
- ✅ chatCmd with full chat loop, streaming, and commands
- ✅ models list with tabwriter formatted output
- ✅ API error handling with errors.As
- ✅ --model persistent flag with config fallback

- Key cobra concepts covered in phase 4 https://docs.google.com/document/d/17R70SuUJC8udPt0YT5auKJo2Fo6OgAQgi8OvUmDM0r0/edit?tab=t.0#heading=h.7bcx6zhlmoxq

```go
```