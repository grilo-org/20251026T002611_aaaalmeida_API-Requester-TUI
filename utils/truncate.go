package utils

func Truncate(s string, maxWidth int) string {
	if len(s) <= maxWidth {
		return s
	}
	return s[:maxWidth-3] + "..."
}
