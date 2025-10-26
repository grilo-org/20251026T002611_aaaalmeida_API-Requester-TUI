package request

import (
	"fmt"
)

type Request struct {
	ID                   int
	Name                 string
	Url                  string
	Method_id            int
	Collection_id        int
	Expected_Status_code *int
	Headers              map[string]string
	Body                 []byte
	BodyType             BodyType
	// TODO: talvez mudar para Date
	Created_at string
	Updated_at string
}

func (r Request) String() string {
	return fmt.Sprintf("%p %d %s %s %d %d %d %s %s %d %s %s",
		&r, r.ID, r.Name, r.Url, r.Method_id, r.Collection_id, &r.Expected_Status_code, r.Headers, r.Body, r.BodyType, r.Created_at, r.Updated_at)
}
