package selectable

type Selectable interface {
	Label() string
	Value() any
}
