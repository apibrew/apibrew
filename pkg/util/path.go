package util

import (
	"strings"
	"unicode"
)

func PathSlug(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			builder.WriteRune('-')
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	return builder.String()
}
