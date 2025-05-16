#!/bin/bash
_akira_complete() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    local suggestions
    suggestions=$(akira suggest "${COMP_LINE}")
    COMPREPLY=($(compgen -W "$suggestions" -- "$cur"))
}
complete -F _akira_complete akira
