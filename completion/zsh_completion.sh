#compdef akira

_akira_complete() {
    local suggestions
    suggestions=("${(@f)$(akira suggest $BUFFER)}")
    compadd -- $suggestions
}

compdef _akira_complete akira
