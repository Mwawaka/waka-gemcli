package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type History struct {
	historyPath string
}

func NewHistory() (*History, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, fmt.Errorf("could not find home directory: %w", err)
	}

	path := filepath.Join(homeDir, ".local", "share", "gemcli", "history.json")

	return &History{
		historyPath: path,
	}, nil
}

func (h *History) SaveHistory(target any) error {
	dir := filepath.Dir(h.historyPath)

	// Hardened 0755 - 0700
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating application data directory: %w", err)
	}

	// Hardened 0644 - 0600
	file, err := os.OpenFile(h.historyPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)

	if err != nil {
		return err
	}

	defer file.Close()

	if err := json.NewEncoder(file).Encode(target); err != nil {
		return err
	}

	return nil
}

func (h *History) LoadHistory(target any) error {
	file, err := os.Open(h.historyPath)

	if err != nil {
		return err
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(target); err != nil {
		return err
	}

	return nil
}
