package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

type History struct {
	HistoryPath string
}

func NewHistory() (*History, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, fmt.Errorf("could not find home directory: %w", err)
	}

	path := filepath.Join(homeDir, ".local", "share", "gemcli", "history.json")

	return &History{
		HistoryPath: path,
	}, nil
}

func (h *History) SaveHistory(chatHistory []*genai.Content) error {
	dir := filepath.Dir(h.HistoryPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating application data directory: %w", err)
	}

	file, err := os.OpenFile(h.HistoryPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	if err := json.NewEncoder(file).Encode(chatHistory); err != nil {
		return err
	}

	return nil
}

func (h *History) LoadHistory() ([]*genai.Content, error) {
	var chatHistory []*genai.Content
	file, err := os.Open(h.HistoryPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&chatHistory); err != nil {
		return nil, err
	}

	return chatHistory, nil
}
