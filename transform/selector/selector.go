package selector

import (
	"strings"

	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/transform"
	"github.com/YuShuanHsieh/trello-transform/utilities"
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
	if !utilities.IsEmptyString(compare.Name) && !strings.Contains(list.Name, compare.Name) {
		return false
	}
	return compare.Closed == list.Closed
}
