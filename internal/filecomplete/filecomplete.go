package filecomplete

import (
	"os"
)

func ListFilesInDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files, nil
}
