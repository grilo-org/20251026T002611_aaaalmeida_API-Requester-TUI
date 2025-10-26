package utils

// remove anything from an array based on index
func RemoveFromArray[Type any](s []Type, i int) []Type {
	return append(s[:i], s[i+1:]...)
}
