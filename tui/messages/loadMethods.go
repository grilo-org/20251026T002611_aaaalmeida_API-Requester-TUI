package messages

import "api-requester/domain/method"

type LoadMethodsMsg struct {
	Methods []method.Method
	Err     error
}
