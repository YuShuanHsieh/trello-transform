package transform

import (
	"github.com/YuShuanHsieh/trello-transform/models"
)

type SeletorFunc func(*models.Card) bool

func (t *Transform) SelectByList(l *models.List) SeletorFunc {
	if l == nil {
		return nil
	}
	return func(c *models.Card) bool {
		value, ok := t.listsMap[c.IDList]
		match := true
		if ok {
			match = match && l.Closed == value.Closed
		}
		if l.ID != "" {
			match = match && l.ID == c.IDList
		}
		if l.Name != "" && ok {
			match = match && l.Name == value.Name
		}
		return match
	}
}
