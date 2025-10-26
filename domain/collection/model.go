package collection

import (
	"api-requester/domain/request"
	"database/sql"
	"fmt"
)

type Collection struct {
	ID          int
	Name        string
	Requests    []*request.Request
	Created_at  string
	Updated_at  string
	Description sql.NullString
}

func (c Collection) String() string {
	return fmt.Sprintf("%d %s %s", c.ID, c.Name, c.Description.String)
}
