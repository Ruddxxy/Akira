package filecomplete

import (
    "io/ioutil"
    "os"
)

func ListFiles(path string) []string {
    var files []string
    fi, err := os.Stat(path)
    if err != nil || !fi.IsDir() {
        return files
    }

    entries, err := ioutil.ReadDir(path)
    if err != nil {
        return files
    }

    for _, e := range entries {
        files = append(files, e.Name())
    }
    return files
}
