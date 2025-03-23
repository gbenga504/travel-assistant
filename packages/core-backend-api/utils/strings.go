package utils

import "strings"

func FirstLetterToLower(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[0:1]) + s[1:]
}
