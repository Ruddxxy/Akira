package alias

var Aliases = map[string]string{
	"gcm": "git commit -m",
	"gst": "git status",
}

func Resolve(alias string) string {
	if val, ok := Aliases[alias]; ok {
		return val
	}
	return alias
}
