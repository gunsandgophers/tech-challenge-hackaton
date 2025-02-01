package utils

import "strings"

func AssertIn[T comparable](value T, list []T) bool {
	for _, element := range list {
		if element == value {
			return true
		}
	}
	return false
}

func AssertNotEmpty(value string) bool {
	return strings.TrimSpace(value) != ""
}
