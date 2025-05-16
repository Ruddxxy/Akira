package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DiscoverPathCommands() []string {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	cmds := make(map[string]struct{})

	for _, dir := range paths {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			if !file.IsDir() {
				full := filepath.Join(dir, file.Name())
				info, err := os.Stat(full)
				if err == nil && info.Mode()&0111 != 0 {
					cmds[file.Name()] = struct{}{}
				}
			}
		}
	}
	list := []string{}
	for k := range cmds {
		list = append(list, k)
	}
	return list
}

func InstallCompletionScript(shell string) {
	switch shell {
	case "bash":
		fmt.Println("source ~/.akira/scripts/akira_completion.sh")
	case "zsh":
		fmt.Println("compdef _akira akira")
	default:
		fmt.Println("Shell not supported")
	}
}
