#!/bin/bash
_akira_complete() {
    COMPREPLY=($(akira suggest "${COMP_LINE}" | compgen -W "$(_akira_suggest_words)" -- "${COMP_WORDS[COMP_CWORD]}"))
}

_akira_suggest_words() {
  akira suggest "${COMP_LINE}"
}

complete -F _akira_complete akira
