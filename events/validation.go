package events

import (
	"regexp"
)

func ValidateTitle(title string) bool {
	pattern := "^[a-zA-Z0-9 ]{3,50}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
