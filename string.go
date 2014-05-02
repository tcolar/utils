// History: May 02 14 tcolar Creation

package utils

import (
	"strings"
	"unicode"
)

// IsEmptyString tells us if a string is empty or not
func IsEmptyString(str string) bool {
	return len(str) == 0
}

// UnderscoredToSnake converts an underscored name to a snake case name
// ie "this_is_an_example" -> "ThisIsAnExample"
func UnderscoredToSnake(str string) string {
	snaked := []string{""}
	for _, part := range strings.Split(str, "_") {
		snaked = append(snaked, strings.Title(part))
	}
	return strings.Join(snaked, "")
}

// SnakeToUnderscore converts a snake case string to a underscored name
// ie "ThisIsAnExample" -> "this_is_an_example"
func SnakeToUnderscore(str string) string {
	result := []rune{}
	for _, c := range str {
		if len(result) > 0 && unicode.IsUpper(c) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(c))
	}
	return string(result)
}
