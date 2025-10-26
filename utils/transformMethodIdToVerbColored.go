package utils

import "net/http"

func TransformMethodIdToVerbColored(id int) string {
	var Reset = "\033[0m"

	var Black = "\033[30m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Cyan = "\033[36m"
	var Gray = "\033[37m"

	var verb string
	switch id {
	case 1:
		verb = Green + http.MethodGet
	case 2:
		verb = Blue + http.MethodPost
	case 3:
		verb = Yellow + http.MethodPut
	case 4:
		verb = Red + http.MethodDelete
	case 5:
		verb = Magenta + http.MethodPatch
	case 6:
		verb = Cyan + http.MethodHead
	case 7:
		verb = Gray + http.MethodTrace
	case 8:
		verb = Black + http.MethodOptions
	}
	verb += Reset
	return verb
}
