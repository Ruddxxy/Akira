package plugins

import (
	"fmt"
	"os"
)

var pluginMap = map[string]string{}

func RegisterPlugin(name string, scriptPath string) error {
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script does not exist")
	}
	pluginMap[name] = scriptPath
	return nil
}

func GetPlugin(name string) string {
	return pluginMap[name]
}
