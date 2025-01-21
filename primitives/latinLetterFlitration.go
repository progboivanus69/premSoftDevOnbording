package primitives

import (
	"strings"
	"unicode"
)

func LatinLetters(s string) string {
	f := ""
	sb := &strings.Builder{}
	for _, val := range s {
		if unicode.Is(unicode.Latin, val) {
			sb.WriteString(string(val))
			f = sb.String()
		}
	}
	return f
}
