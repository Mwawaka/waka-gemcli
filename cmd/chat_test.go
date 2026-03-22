package cmd

import (
	"context"
	"iter"

	"google.golang.org/genai"
)


type MockChatSession struct {
	response *genai.GenerateContentResponse
	err      error
}

func (m *MockChatSession) SendMessageStream(ctx context.Context, parts ...genai.Part) iter.Seq2[*genai.GenerateContentResponse, error] {
	return func(yield func(*genai.GenerateContentResponse, error) bool) {
		if m.err != nil {
			yield(nil, m.err)
			return
		}
		yield(m.response, nil)
	}
}
