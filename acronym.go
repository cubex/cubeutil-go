package cubeutil

import "strings"

// Acronym forms an acronym from an input string
func Acronym(source string, length int) string {
	nameParts := strings.Split(source, " ")
	if length == 0 {
		length = len(nameParts)
	}
	acronym := ""
	for i, part := range nameParts {
		for x := 0; x < length; x++ {
			acronym += string([]rune(part)[x])
			if i+1 < len(nameParts) || len(acronym) == length {
				break
			}
		}
		if len(acronym) == length {
			break
		}
	}
	return strings.ToUpper(acronym)
}
