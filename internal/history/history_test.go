package history

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

type testData struct {
	APIKey string `json:"apiKey"`
	Model  string `json:"model"`
}

func Test_SaveHistory(t *testing.T) {
	history := newHistoryWithCustomPath(t.TempDir())

	if err := history.SaveHistory(&testData{
		APIKey: "apiKey",
		Model:  "model",
	}); err != nil {
		t.Fatal(err)
	}
}

func Test_LoadHistory(t *testing.T) {
	var testResult testData
	testInput := testData{
		APIKey: "apiKey",
		Model:  "model",
	}
	history := newHistoryWithCustomPath(t.TempDir())

	if err := history.SaveHistory(&testInput); err != nil {
		t.Fatal(err)
	}

	if err := history.LoadHistory(&testResult); err != nil {
		t.Fatal(err)
	}

	if testResult.APIKey != testInput.APIKey {
		t.Errorf("expected %q, got %q", testInput.APIKey, testResult.APIKey)
	}

	if testResult.Model != testInput.Model {
		t.Errorf("expected %q, got %q", testInput.Model, testResult.Model)
	}

}

func Test_LoadHistory_FileNotExist(t *testing.T) {
	var testResult testData
	history := newHistoryWithCustomPath(t.TempDir())

	if err := history.LoadHistory(&testResult); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return
		} else {
			t.Fatal(err)
		}
	}

	t.Fatal("expected an error, got nil")
}

func newHistoryWithCustomPath(dir string) *History {
	path := filepath.Join(dir, ".local", "share", "gemcli", "history.json")

	return &History{
		historyPath: path,
	}
}
