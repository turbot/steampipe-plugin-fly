package errors

import (
	"regexp"
)

func NotFoundError(err error) bool {
	notFoundErr := "(?i)Could not find"
	expectedErr := regexp.MustCompile(notFoundErr)
	return expectedErr.Match([]byte(err.Error()))
}
