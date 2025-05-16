package alias

import (
    "bufio"
    "os"
    "regexp"
    "strings"
    "path/filepath"
)

func LoadAliases() (map[string]string, error) {
    aliases := make(map[string]string)
    home, err := os.UserHomeDir()
    if err != nil {
        return aliases, err
    }

    files := []string{
        filepath.Join(home, ".bashrc"),
        filepath.Join(home, ".zshrc"),
    }

    for _, file := range files {
        f, err := os.Open(file)
        if err != nil {
            continue
        }
        defer f.Close()

        scanner := bufio.NewScanner(f)
        re := regexp.MustCompile(`alias\s+(\w+)='([^']+)'`)
        for scanner.Scan() {
            line := scanner.Text()
            matches := re.FindStringSubmatch(line)
            if len(matches) == 3 {
                aliasName := matches[1]
                aliasCmd := matches[2]
                aliases[aliasName] = aliasCmd
            }
        }
    }
    return aliases, nil
}
