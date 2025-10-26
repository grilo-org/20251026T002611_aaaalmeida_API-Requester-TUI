package messages

import "api-requester/domain/request"

type SendRequestMsg struct {
	Request *request.Request
	Err     error
}
