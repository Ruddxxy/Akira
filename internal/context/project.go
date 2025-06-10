package context

import (
	"os"
)

func InGitProject() bool {
	_, err := os.Stat(".git")
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
