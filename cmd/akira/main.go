package main

import (
	"akira/internal/alias"
	"akira/internal/filecomplete"
	"akira/internal/gui"
	"akira/internal/plugins"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Akira: Intelligent CLI Autocomplete And Discovery")
		fmt.Println("Usage: akira [command]")
		fmt.Println("Commands:")
		fmt.Println((" register-plugin [name][script path]     Register a plugin with a name and script path"))
		fmt.Println("  suggest [filter]       - Suggest commands based on filter")
		fmt.Println("  install-completion [shell] - Install completion script for shell (bash, zsh)")
		fmt.Println("  register-plugin [name] [script_path] - Register a plugin with a name and script path")
		fmt.Println(" list-plugins         - List all registered plugins")
		fmt.Println("  get-aliases           - Get all user-defined aliases")
		fmt.Println("  help                   - Show this help message")
		fmt.Println("  version                - Show version information")
		fmt.Println("  exit                   - Exit the program")
		fmt.Println("  quit                   - Quit the program")

		os.Exit(0)
	}

	home, _ := os.UserHomeDir()
	if err := alias.LoadUserAliases(home); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading aliases: %v\n", err)
	}
	if err := plugins.LoadUserPlugin(home); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading plugins: %v\n", err)
	}

	switch args[0] {
	case "gui":
		gui.RunGUI()
	case "suggest":
		filter := ""
		if len(args) > 1 {
			filter = args[1] // Optional filter argument
		}
		for _, cmd := range filecomplete.GetSuggestions(filter) {
			fmt.Println(cmd)
		}
	case "install-completion":
		if len(args) < 2 {
			fmt.Println("Specify Shell For Installation: bash, zsh")
			return
		}
	case "register-plugin":
		if len(args) < 3 {
			fmt.Println("Usage: akira register-plugin [name] [script_path]")
			return
		}
		err := plugins.RegisterPlugin(args[1], args[2])
		if err != nil {
			fmt.Printf("Error registering plugin: %v\n", err)
		}
	default:
		fmt.Println("Unknown command:", args[0])
	}
}
