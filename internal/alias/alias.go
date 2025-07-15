package alias

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

var (
	userAliases = make(map[string]map[string]string)
	aliases     = map[string]map[string]string{}
	mutex       sync.RWMutex
)

// Aliases holds the default aliases.
func LoadUserAliases(home string) error {
	path := filepath.Join(home, ".akira_aliases.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//file does not exist, create it with default values
		defaults := map[string]map[string]string{
			"gh": {
				"status":   "git status",
				"commit":   "git commit -m",
				"push":     "git push",
				"pull":     "git pull",
				"clone":    "git clone",
				"branch":   "git branch",
				"checkout": "git checkout",
				"merge":    "git merge",
				"rebase":   "git rebase",
				"reset":    "git reset",
				"log":      "git log",
				"diff":     "git diff",
				"stash":    "git stash",
				"tag":      "git tag",
				"remote":   "git remote"},
			"dc": {
				"build":   "docker build",
				"run":     "docker run ",
				"compose": "docker-compose",
				"ps":      "docker ps",
				"logs":    "docker logs",
				"exec":    "docker exec -it",
				"images":  "docker images",
				"rmi":     "docker rmi",
				"rm":      "docker rm",
				"network": "docker network",
				"volume":  "docker volume",
			},
			"os": {
				"ls":    "ls -la",
				"cd":    "cd",
				"mkdir": "mkdir",
				"rm":    "rm -rf",
				"echo":  "echo",
				"cat":   "cat",
				"touch": "touch",
				"cp":    "cp -r",
				"mv":    "mv",
				"pwd":   "pwd"},
			"py": {
				"run":        "python",
				"install":    "pip install",
				"freeze":     "pip freeze",
				"venv":       "python -m venv",
				"activate":   "source venv/bin/activate",
				"deactivate": "deactivate",
				"test":       "python -m unittest",
			},
			"js": {
				"run":              "node",
				"install":          "npm install",
				"start":            "npm start",
				"test":             "npm test",
				"build":            "npm run build",
				"init":             "npm init -y",
				"serve":            "npx serve",
				"create-react-app": "npx create-react-app",
			},
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(defaults); err != nil {
			return err
		}
	}
	//Load the aliases from the file
	return loadfromFile(path)
}
func loadfromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	var userAliases = make(map[string]map[string]string)
	if err := json.NewDecoder(file).Decode(&userAliases); err != nil {
		return err
	}
	mutex.Lock()

	defer mutex.Unlock()
	aliases = userAliases
	return nil
}
func GetAliases() map[string]map[string]string {
	mutex.RLock()
	defer mutex.RUnlock()
	return aliases
}
func SaveAliases(home string) error {
	path := filepath.Join(home, ".akira_aliases.json")
	mutex.RLock()
	defer mutex.RUnlock()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(userAliases)
}

// Resolve returns the resolved alias or the original command if not found.
func Resolve(group, key string) (string, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	if cmds, ok := aliases[group]; ok {
		if cmd, ok := cmds[key]; ok {
			return cmd, true
		}
	}
	return "", false
}
