package plugins

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

var (
	pluginMap   = make(map[string]string)
	pluginMutex sync.RWMutex
)

func RegisterPlugin(name string, scriptPath string) error {
	if _, err := os.Stat(scriptPath); err != nil {
		return fmt.Errorf("cannot access script: %w", err)
	}
	pluginMutex.Lock()
	defer pluginMutex.Unlock()
	pluginMap[name] = scriptPath
	return nil
}

func GetPlugin(name string) (string, error) {
	pluginMutex.RLock()
	defer pluginMutex.RUnlock()
	scriptPath, ok := pluginMap[name]
	if !ok {
		return "", errors.New("plugin not found")
	}
	return scriptPath, nil
}
