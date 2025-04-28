package utils

import (
	"regexp"
	"strings"
)

// ValidateText checks if the input text meets certain criteria.
func ValidateText(text string) bool {
	// Example validation: text should not be empty and should not exceed 500 characters
	if len(text) == 0 || len(text) > 500 {
		return false
	}
	return true
}

// FormatText trims whitespace and converts text to lowercase.
func FormatText(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

// IsValidUserID checks if the user ID is in a valid format (e.g., alphanumeric).
func IsValidUserID(userID string) bool {
	// Example regex for alphanumeric user ID
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(userID)
}