package commands

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func DiscoverPathCommands() ([]string, []error) {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	cmds := make(map[string]struct{})
	var errs []error

	for _, dir := range paths {
		files, err := os.ReadDir(dir)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading directory %s: %w", dir, err))
			continue
		}

		for _, file := range files {
			if !file.IsDir() {
				name := file.Name()
				if isExecutable(name) {
					cmds[stripExt(name)] = struct{}{}
				}
			}
		}
	}
	var list []string
	for k := range cmds {
		list = append(list, k)
	}
	return list, errs
}

func isExecutable(name string) bool {
	if runtime.GOOS == "windows" {
		return strings.HasSuffix(strings.ToLower(name), ".exe") ||
			strings.HasSuffix(strings.ToLower(name), ".bat") ||
			strings.HasSuffix(strings.ToLower(name), ".cmd")
	}
	return true
}

func stripExt(name string) string {
	if runtime.GOOS == "windows" {
		name = strings.TrimSuffix(name, ".exe")
		name = strings.TrimSuffix(name, ".bat")
		name = strings.TrimSuffix(name, ".cmd")

	}
	return name
}
func InstallCompletionScript(shell string) error {
	switch shell {
	case "bash":
		fmt.Println("source ~/.akira/completion.bash")
	case "zsh":
		fmt.Println("compdef _akira akira")
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}
	return nil
}
