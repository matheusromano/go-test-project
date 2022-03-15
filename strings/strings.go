package strings

import "strings"

func StringCompare(a, b string) int {
	result := strings.Compare(a, b)
	return result
}

func StringContains(completeString, substring string) bool {
	result := strings.Contains(completeString, substring)
	return result
}
