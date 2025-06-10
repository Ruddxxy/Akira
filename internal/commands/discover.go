package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DiscoverPathCommands() ([]string, []error) {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	cmds := make(map[string]struct{})

	var lastErr error
	readableDirFound := false
	dirErrors := make([]error, 0)

	for _, dir := range paths {
		files, err := os.ReadDir(dir)
		if err != nil {
			dirErrors = append(dirErrors, fmt.Errorf("error reading directory %s: %w", dir, err))
			lastErr = err
			continue
		}
		readableDirFound = true
		for _, file := range files {
			if !file.IsDir() {
				full := filepath.Join(dir, file.Name())
				info, err := os.Stat(full)
				if err != nil {
					dirErrors = append(dirErrors, fmt.Errorf("error getting file info for %s: %w", full, err))
					continue
				}
				if info.Mode()&0111 != 0 {
					cmds[file.Name()] = struct{}{}
				}
			}
		}
	}
	list := []string{}
	for k := range cmds {
		list = append(list, k)
	}
	if !readableDirFound && lastErr != nil {
		dirErrors = append(dirErrors, fmt.Errorf("failed to read any PATH directories: last error: %w", lastErr))
	}
	return list, dirErrors
}

func InstallCompletionScript(shell string) error {
	switch shell {
	case "bash":
		fmt.Println("source ~/.akira/scripts/akira_completion.sh")
	case "zsh":
		fmt.Println("compdef _akira akira")
	default:
		return fmt.Errorf("shell not supported: %s", shell)
	}
	return nil
}
