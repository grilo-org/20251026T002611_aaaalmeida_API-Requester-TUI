package messages

import "api-requester/domain/collection"

type LoadCollectionsMsg struct {
	Collections []*collection.Collection
	Err         error
}
