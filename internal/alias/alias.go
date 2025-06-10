package alias

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

var Aliases = map[string]string{
	"gcm": "git commit -m",
	"gst": "git status",
}

// LoadUserAliases loads aliases from a JSON file and merges them.

var aliasesMutex sync.RWMutex

func LoadUserAliases(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	userAliases := make(map[string]string)
	if err := json.NewDecoder(file).Decode(&userAliases); err != nil {
		return err
	}

	aliasesMutex.Lock()
	defer aliasesMutex.Unlock()
	for k, v := range userAliases {
		if _, exists := Aliases[k]; exists {
			// Log a warning about overwriting an existing alias
			// You can replace this with a logger if desired
			os.Stderr.WriteString("Warning: overwriting existing alias '" + k + "'\n")
		}
		Aliases[k] = v
	}
	return nil
}

func Resolve(alias string) string {
	aliasesMutex.RLock()
	defer aliasesMutex.RUnlock()
	if val, ok := Aliases[alias]; ok {
		return val
	}
	return alias
}

// ExpandAlias recursively expands an alias until it is fully resolved.
func ExpandAlias(command string) string {
	return expandAliasWithDepth(command, make(map[string]bool), 0)
}

func expandAliasWithDepth(command string, visited map[string]bool, depth int) string {
	const maxDepth = 10
	if depth > maxDepth {
		return command // Prevent deep recursion
	}

	parts := strings.SplitN(command, " ", 2)
	alias := parts[0]

	if visited[alias] {
		return command // Circular reference detected
	}

	args := ""
	if len(parts) > 1 {
		args = " " + parts[1]
	}

	resolved := Resolve(alias)
	if resolved == alias {
		return command
	}

	visited[alias] = true
	return expandAliasWithDepth(resolved+args, visited, depth+1)
}
