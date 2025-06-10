package filecomplete

import (
	"os"
	"path/filepath"
)

func ListFilesInDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		files = append(files, filepath.Join(dir, entry.Name()))
	}
	return files, nil
}
