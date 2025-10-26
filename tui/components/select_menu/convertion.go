package select_menu

import (
	"api-requester/domain/collection"
	"api-requester/domain/method"
	"api-requester/shared/selectable"
)

// Conversors
func ConvertMethodsToSelectOptions(methods []method.Method) []selectable.Selectable {
	opts := make([]selectable.Selectable, len(methods))
	for i, m := range methods {
		opts[i] = &MethodOption{m}
	}
	return opts
}

func CollectionsToselectOptions(cols []collection.Collection) []selectable.Selectable {
	opts := make([]selectable.Selectable, len(cols))
	for i, c := range cols {
		opts[i] = CollectionOption{c}
	}
	return opts
}

// Wrappers
type MethodOption struct{ method.Method }

func (o *MethodOption) Label() string { return o.Method.Name }
func (o *MethodOption) Value() any    { return o.Method.ID }

type CollectionOption struct{ collection.Collection }

func (o CollectionOption) Label() string { return o.Collection.Name }
func (o CollectionOption) Value() any    { return o.Collection.ID }
