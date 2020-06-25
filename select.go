package transform

import (
	"github.com/adlio/trello"
)

func SelectByListNames(names ...string) SelectFunc {
	m := make(map[string]struct{})
	for _, name := range names {
		if name != "" {
			m[name] = struct{}{}
		}
	}
	return func(card *trello.Card) bool {
		if card.List == nil {
			return false
		}
		_, ok := m[card.List.Name]
		return ok
	}
}
