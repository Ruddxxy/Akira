package plugins

import (
	"errors"
	"fmt"
	"os"
)

var pluginMap = map[string]string{}

func RegisterPlugin(name string, scriptPath string) error {
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script does not exist: %w", err)
	}
	pluginMap[name] = scriptPath
	return nil
}

func GetPlugin(name string) (string, error) {
	scriptPath, ok := pluginMap[name]
	if !ok {
		return "", errors.New("plugin not found")
	}
	return scriptPath, nil
}
