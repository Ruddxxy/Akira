package context

import (
	"os"
)

func InGitProject() bool {
	_, err := os.Stat(".git")
	return err == nil
}
