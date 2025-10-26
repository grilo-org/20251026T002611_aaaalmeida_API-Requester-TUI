package messages

import "api-requester/domain/request"

type LoadRequestMsg struct {
	Request *request.Request
	Err     error
}
