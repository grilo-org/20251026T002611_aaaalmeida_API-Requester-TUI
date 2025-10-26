package collection_menu

import (
	"api-requester/context"
	"api-requester/domain/collection"
)

const WIDTH int = 30
const HEIGHT int = 25
const PADDING int = 0

// TODO: ajust based on config dotfile
type model struct {
	context        *context.AppContext
	collections    []*collection.Collection
	openCloseIndex []bool
	cursor         cursor
}

type cursor struct {
	colIndex int
	reqIndex *int // if nil then points to collection, else points to requests
}

func NewModel(ctx *context.AppContext) model {
	return model{
		context:        ctx,
		collections:    nil,
		openCloseIndex: nil,
		cursor:         cursor{colIndex: 0, reqIndex: nil},
	}
}
