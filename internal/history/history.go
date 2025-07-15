package history

import (
	"bufio"
	"os"
	"os/user"
	"path/filepath"
)

func LoadHistory(maxEntries int) ([]string, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	historyFile := filepath.Join(user.HomeDir, ".bash_history")
	file, err := os.Open(historyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) > maxEntries {
		lines = lines[len(lines)-maxEntries:] // Keep only the last maxEntries

	}
	return lines, scanner.Err()
}
