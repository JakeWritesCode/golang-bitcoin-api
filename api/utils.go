package api

import (
	"strings"
)

// An implementation similar to pythons f-string in golang. Should I be doing this? We may never know...
func fstring(format string, args ...string) string {
	for i, v := range args {
		if i%2 == 0 {
			args[i] = "{" + v + "}"
		}
	}
	r := strings.NewReplacer(args...)
	return r.Replace(format)
}
