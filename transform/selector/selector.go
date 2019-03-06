package selector

import (
	"strings"

	"github.com/YuShuanHsieh/trello-transform/helpers/validators"
	"github.com/YuShuanHsieh/trello-transform/transform"
	"github.com/YuShuanHsieh/trello-transform/trello"
)

func ByList(compare trello.List) transform.Seletor {
	return func(ctx *transform.Context, c *trello.Card) bool {
		list, ok := ctx.Lists[c.IDList]
		if !ok {
			return false
		}
		return compareList(list, compare)
	}
}

func compareList(list trello.List, compare trello.List) bool {
	if !validators.IsEmptyString(compare.Name) && !strings.Contains(list.Name, compare.Name) {
		return false
	}
	return compare.Closed == list.Closed
}
