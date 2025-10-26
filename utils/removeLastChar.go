package utils

func RemoveLastChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1]
}
