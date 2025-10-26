package utils

import "strings"

func Concatenate(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	var resp strings.Builder
	for _, s := range args {
		resp.WriteString(" " + s)
	}
	return resp.String()
}
