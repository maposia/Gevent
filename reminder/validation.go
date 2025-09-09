package reminder

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidateMessage(message string) error {

	pattern := "^[a-zA-Z0-9 ]{3,50}$"
	matched, err := regexp.MatchString(pattern, message)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New(fmt.Sprintf("message contains invalid characters or is empty"))
	}

	return nil
}
