## 🚀 Akira - Intelligent CLI & GUI Tool

Akira is a modular CLI & GUI tool to help beginners and power users manage command-line aliases, auto-complete suggestions, and plugins. It’s designed to make working with Git, Docker, Linux, Node, Python and more super easy — right from your terminal or with a friendly GUI.

---

### ✨ Features

✅ Nested JSON alias groups (`gh`, `dc`, `os`, `py`, `np`)
✅ Auto-creates `~/.akira_aliases.json` with smart defaults on first run
✅ CLI commands for suggestions & alias resolution
✅ Fyne GUI to view, add, and search aliases interactively
✅ Shell completions for Bash, Zsh, and PowerShell
✅ Modular Go project structure with internal packages

---

### 📂 Project structure

```
akira/
├── cmd/akira/main.go         # CLI entry
├── internal/
│   ├── alias/                # Nested alias loader, saver, resolver
│   ├── filecomplete/         # Suggestion builder
│   ├── plugins/              # Plugin loader
│   ├── discover/             # Path discovery
│   └── gui/                  # Fyne GUI
├── scripts/
│   ├── akira_completion.sh
│   ├── akira_completion.zsh
│   └── akira_completion.ps1
├── go.mod
├── go.sum
└── README.md
```

---

### 🚀 Getting started

#### 🔧 Build

**Linux / Mac:**

```bash
go build -o akira ./cmd/akira
```

**Windows:**

```powershell
go build -ldflags "-H windows" -o akira.exe ./cmd/akira
```

**Install globally:**

```bash
go install ./cmd/akira
```

---

### ⚙️ Usage

#### 🚀 CLI commands

```bash
akira suggest gh
# gh status -> git status
# gh commit -> git commit

akira run-alias gh status
# Resolved to: git status
```

#### 🎨 GUI

```bash
akira gui
```

* View all aliases by group
* Add new aliases visually
* Try out live suggestions

---

### 🐚 Shell completions

**Bash:**

```bash
source scripts/akira_completion.sh
```

**Zsh:**

```bash
source scripts/akira_completion.zsh
```

**PowerShell:**

```powershell
. scripts/akira_completion.ps1
```

✅ Now `akira [TAB]` completes your commands.

---

### 🚀 Testing checklist

| Test Case              | Command                     | Expected Result                  |
| ---------------------- | --------------------------- | -------------------------------- |
| JSON auto-creation     | `akira suggest`             | Creates `~/.akira_aliases.json`  |
| Suggest nested group   | `akira suggest gh`          | Lists only git commands          |
| Run alias resolution   | `akira run-alias gh status` | Prints `Resolved to: git status` |
| GUI launches           | `akira gui`                 | Opens Fyne window                |
| Shell completions work | `akira [TAB]`               | Shows completions                |

---

### 💾 Data files

* User aliases are stored in:

| OS        | File                                 |
| --------- | ------------------------------------ |
| Linux/Mac | `~/.akira_aliases.json`              |
| Windows   | `C:\Users\<You>\.akira_aliases.json` |

Created with defaults on first run.

---

### ❤️ Contributing

PRs and feature suggestions are welcome!

---

### 📝 License

MIT
