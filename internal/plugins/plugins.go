package plugins

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	pluginMap = make(map[string]string)
	mux       sync.RWMutex
)

func LoadUserPlugin(home string) error {
	path := filepath.Join(home, ".akira", "plugins.json")
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, we can ignore this error
			return nil
		}
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&pluginMap)
}

func RegisterPlugin(name, scriptpath string) error {
	if _, err := os.Stat(scriptpath); err != nil {
		return fmt.Errorf("cannot access script path: %w", err)
	}
	mux.Lock()
	pluginMap[name] = scriptpath
	mux.Unlock()
	return SavePlugins()
}

func SavePlugins() error {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".akira", "plugins.json")
	mux.RLock()
	defer mux.RUnlock()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(pluginMap)
}

func GetPlugins() map[string]string {
	mux.RLock()
	defer mux.RUnlock()
	cp := make(map[string]string, len(pluginMap))
	for k, v := range pluginMap {
		cp[k] = v
	}
	return cp
}
