package history

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func LoadHistory() []string {
	user, err := user.Current()
	if err != nil {
		return []string{}
	}
	historyFile := filepath.Join(user.HomeDir, ".bash_history")
	data, err := os.ReadFile(historyFile)
	if err != nil {
		return []string{}
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
