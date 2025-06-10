#!/bin/bash
_akira_complete() {
    local suggestions
    suggestions=$(akira suggest "${COMP_LINE}")
    COMPREPLY=( $(compgen -W "$suggestions" -- "${COMP_WORDS[COMP_CWORD]}") )
}

complete -F _akira_complete akira
