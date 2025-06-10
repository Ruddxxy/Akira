#compdef akira

_akira_complete() {
    local -a suggestions
    suggestions=("${(@f)$(akira suggest "$BUFFER")}")
    _describe 'completion' suggestions
}

compdef _akira_complete akira
