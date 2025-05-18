#!/bin/bash
_akira_completions() {
    COMPREPLY=($(akira suggest | grep "^${COMP_WORDS[1]}"))
}
complete -F _akira_completions akira
