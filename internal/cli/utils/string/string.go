package string

import "strings"

// ToUpperTrim trims spaces and converts string to uppercase
func ToUpperTrim(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

// IsYes returns true if input string looks like a "yes"
func IsYes(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == "yes" || s == "y" || s == "true" || s == "1"
}
