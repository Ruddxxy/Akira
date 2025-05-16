package main

import (
	"akira/internal/commands"
	"akira/internal/history"
	"akira/internal/plugins"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Akira: Intelligent CLI autocomplete and discovery")
		fmt.Println("Usage: akira [command]")
		os.Exit(0)
	}

	switch args[0] {
	case "suggest":
		all := commands.DiscoverPathCommands()
		historyItems := history.LoadHistory()
		all = append(all, historyItems...)
		for _, cmd := range all {
			fmt.Println(cmd)
		}
	case "install-completion":
		if len(args) < 2 {
			fmt.Println("Specify shell: bash, zsh, powershell")
			return
		}
		commands.InstallCompletionScript(args[1])
	case "register-plugin":
		if len(args) < 3 {
			fmt.Println("Usage: akira register-plugin [name] [script_path]")
			return
		}
		err := plugins.RegisterPlugin(args[1], args[2])
		if err != nil {
			fmt.Println("Error registering plugin:", err)
		}
	default:
		fmt.Println("Unknown command:", args[0])
	}
}
