# Akira CLI Tool

Akira is a powerful, context-aware CLI autocompletion tool designed to enhance your Bash and Zsh terminal experience by providing smart command, option, and file path suggestions. It improves productivity by learning from your command history, aliases, and dynamically discovered commands, while filtering out private/internal commands.

---

## Features

- **Context-aware suggestions**: Understands the current input context to provide relevant completions.
- **Alias expansion**: Supports user-defined command aliases for faster typing.
- **History-based suggestions**: Leverages your command history for personalized autocomplete.
- **Private command filtering**: Hides internal or private commands (e.g., commands starting with `_`).
- **Shell compatibility**: Supports Bash and Zsh out of the box with easy-to-use completion scripts.
- **Lightweight & fast**: Written in Go for optimal performance.

---

## Prerequisites

- [Go](https://golang.org/dl/) 1.16+ (to build from source)
- Bash or Zsh shell on Linux/macOS or Windows Subsystem for Linux (WSL)
- (Optional) Windows users can test in Linux environments such as Kali Linux or WSL

---

## Installation

### Build from source

Clone the repository and build the executable:

```bash
git clone https://github.com/yourusername/akira.git
cd akira
go mod tidy
go build -o akira ./cmd/akira
```

---

## Usage

After building, you can run Akira with:

```bash
./akira [command] [flags]
```

To enable shell completion, source the provided scripts:

```bash
# For Bash
source completion/akira_completion.sh

# For Zsh
source completion/akira_completion.zsh
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Author

Made by RUDRA MOHAPATRO