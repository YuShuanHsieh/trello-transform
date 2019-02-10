package selector

import (
	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/transform"
)

func ByList(t *transform.Transform, compare models.List) models.SeletorFunc {
	return func(c *models.Card) bool {
		list := t.GetListById(c.IDList)
		if list == nil {
			return false
		}
		return compareList(*list, compare)
	}
}

func compareList(list models.List, compare models.List) bool {
	if compare.Name != "" && compare.Name != list.Name {
		return false
	}
	return compare.Closed == list.Closed
}
