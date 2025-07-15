package filecomplete

import (
	"akira/internal/commands"
	"akira/internal/history"
	"strings"
)

func GetSuggestions(filter string) []string {
	all, _ := commands.DiscoverPathCommands()
	historyItems, _ := history.LoadHistory(100) // Load last 100 history items
	all = append(all, historyItems...)

	//if empty, return all commands
	if len(all) == 0 {
		all = []string{"git status", "git commit", "git push", "ls", "cd", "mkdir", "rm", "echo", "cat", "touch", "cp",
			"mv", "pwd", "python", "node", "npm", "go", "rustc", "cargo", "javac", "java", "gcc", "g++", "make",
			"cmake", "npm", "npm install", "pip", "pip3", "ruby", "rails", "php", "composer", "docker", "docker-compose", "git clone", "git pull", "git fetch", "git branch", "git checkout", "git merge", "git reset", "git log", "git diff", "git status", "sudo", "chmod", "chown", "ps", "top", "htop", "kill", "pkill",
			"df", "du", "free", "uptime", "whoami", "uname", "ifconfig", "ip", "ping", "traceroute",
			"curl", "wget", "scp", "ssh", "ftp", "sftp", "rsync", "tar", "gzip", "gunzip",
			"zip", "unzip", "find", "grep", "sed", "awk", "cut", "sort"}
	}
	var suggestions []string
	for _, cmd := range all {
		if shouldBlacklist(cmd) {
			continue
		}
		if filter == "" || strings.HasPrefix(cmd, filter) {
			suggestions = append(suggestions, cmd)
		}
	}
	return suggestions
}

func shouldBlacklist(cmd string) bool {
	blacklist := []string{"rm", "shutdown", "user.email", "user.name", "token", "password", "secret", "api_key", "access_token", "private_key", "ssh_key", "aws_secret_access_key", "aws_access_key_id",
		"gcp_service_account_key", "azure_client_secret", "azure_client_id", "azure_tenant_id", "db_password", "db_user", "db_host", "db_name", "db_port", "db_connection_string", "ftp_password", "ftp_user",
		"ftp_host", "ftp_port", "smtp_password", "smtp_user", "smtp_host", "smtp_port", "email_password", "email_user", "email_host", "email_port", "api_secret", "api_user", "api_host", "api_port", "api_base_url", "api_endpoint", "api_version", "api_key_id", "api_key_secret"}
	for _, bad := range blacklist {
		if strings.Contains(cmd, bad) {
			return true
		}
	}
	return false
}
