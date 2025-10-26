package method

import "fmt"

type Method struct {
	ID   int
	Name string
}

func (m Method) String() string {
	return fmt.Sprintf("%d %s", m.ID, m.Name)
}
