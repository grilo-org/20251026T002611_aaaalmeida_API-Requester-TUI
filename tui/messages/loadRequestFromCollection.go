package messages

import "api-requester/domain/request"

type LoadRequestFromCollectionMsg struct {
	Collection_id int
	Requests      []*request.Request
	Err           error
}
