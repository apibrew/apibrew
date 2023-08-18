package util

import (
	"strings"
	"unicode"
)

func StripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Pointer[T interface{}](val T) *T {
	var pointer = new(T)

	*pointer = val

	return pointer
}
