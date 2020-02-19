package cmd

import (
	"strings"
	"unicode"
)

func sanitizeStationName(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '-'
		} else if r == '/' {
			return '-'
		}
		return r
	}, str)
}