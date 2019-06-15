package transform

import (
	"strings"

	"github.com/YuShuanHsieh/trello-transform/trello"
)

func ByListNames(names ...string) Seletor {
	m := make(map[string]struct{})
	for _, name := range names {
		m[name] = struct{}{}
	}
	return func(ctx Context, c *trello.Card) bool {
		list, ok := ctx.Lists[c.IDList]
		if ok {
			if _, ok := m[list.Name]; ok {
				return true
			}
		}
		return false
	}
}

func ByList(compare trello.List) Seletor {
	return func(ctx Context, c *trello.Card) bool {
		list, ok := ctx.Lists[c.IDList]
		if !ok {
			return false
		}
		return compareList(list, compare)
	}
}

func compareList(list trello.List, compare trello.List) bool {
	if !isEmptyString(compare.Name) && !strings.Contains(list.Name, compare.Name) {
		return false
	}
	return compare.Closed == list.Closed
}
