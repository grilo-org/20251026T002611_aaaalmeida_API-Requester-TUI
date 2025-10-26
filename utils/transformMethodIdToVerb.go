package utils

import "net/http"

func TransformMethodIdToVerb(id int) string {
	var verb string
	switch id {
	case 1:
		verb = http.MethodGet
	case 2:
		verb = http.MethodPost
	case 3:
		verb = http.MethodPut
	case 4:
		verb = http.MethodDelete
	case 5:
		verb = http.MethodPatch
	case 6:
		verb = http.MethodHead
	case 7:
		verb = http.MethodTrace
	case 8:
		verb = http.MethodOptions
	}
	return verb
}
