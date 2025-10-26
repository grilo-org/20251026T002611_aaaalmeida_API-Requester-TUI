package utils

import "regexp"

var headerValueCharRegex = regexp.MustCompile(`^[\w\d\s\-._:/?#\/'()*+*;,=%]$`)

func IsValidRequestHeaderValue(s string) bool {
	return headerValueCharRegex.MatchString(s)
}
