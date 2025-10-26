package utils

import "regexp"

var urlCharRegex = regexp.MustCompile(`^[A-Za-z0-9\-._~:/?#\[\]@!$&'()*+,;=%]$`)

func IsValidUrlChar(s string) bool {
	return urlCharRegex.MatchString(s)
}
