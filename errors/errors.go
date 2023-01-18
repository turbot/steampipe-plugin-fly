package errors

import (
	"regexp"
)

// NotFoundError: returns true if the queried resource does not exist and the query returns a related error for that.
// If a resource doesn't exist, the query gives below error:
// Error: Could not find ..."
func NotFoundError(err error) bool {
	notFoundErr := "(?i)Could not find"
	expectedErr := regexp.MustCompile(notFoundErr)
	return expectedErr.Match([]byte(err.Error()))
}
