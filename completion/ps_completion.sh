#!/usr/bin/env bash

_complete_cli_tool() {
    local word="${COMP_WORDS[COMP_CWORD]}"
    local prev=""
    if [[ $COMP_CWORD -gt 0 ]]; then
        prev="${COMP_WORDS[COMP_CWORD-1]}"
    fi

    if [[ "$prev" == "cli_tool" ]]; then
        COMPREPLY=( $(compgen -W "command1 command2 command3" -- "$word") )
    elif [[ "$prev" == "command1" ]]; then
        COMPREPLY=( $(compgen -W "--option1 --option2 --option3" -- "$word") )
    elif [[ "$prev" == "command2" ]]; then
        COMPREPLY=( $(compgen -W "argument1 argument2 argument3" -- "$word") )
    else
        COMPREPLY=( $(compgen -W "cli_tool" -- "$word") )
    fi
}

complete -F _complete_cli_tool cli_tool