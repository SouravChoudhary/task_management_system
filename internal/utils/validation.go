package utils

import (
	"errors"
)

// ValidateEnum checks if a value exists in a valid set.
func ValidateEnum(field, value string, valid []string) error {
	for _, v := range valid {
		if v == value {
			return nil
		}
	}
	return errors.New("invalid " + field + ": " + value)
}
