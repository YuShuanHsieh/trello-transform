package transform

import (
	"log"

	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/trello"
)

type Transform struct {
	table        models.TrelloTable
	labelsMap    map[string]*models.Label
	listsMap     map[string]*models.List
	briefFnMap   map[string]ResultFunc
	selectorChan []models.SeletorFunc
	result       map[string]interface{}
}

func New(rawData []byte) *Transform {
	t := initTransform()

	err := trello.UnmarshalJson(rawData, &t.table)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	t.allocateLabelsMap()
	t.allocateListsMap()
	return t
}

func initTransform() *Transform {
	return &Transform{
		listsMap:   make(map[string]*models.List),
		labelsMap:  make(map[string]*models.Label),
		briefFnMap: make(map[string]ResultFunc),
		result:     make(map[string]interface{})}
}

func (t *Transform) allocateLabelsMap() {
	for i, v := range t.table.Labels {
		t.labelsMap[v.ID] = &t.table.Labels[i]
	}
}

func (t *Transform) allocateListsMap() {
	for i, v := range t.table.Lists {
		t.listsMap[v.ID] = &t.table.Lists[i]
	}
}

func (t *Transform) Select(fn models.SeletorFunc) {
	if fn != nil {
		t.selectorChan = append(t.selectorChan, fn)
	}
}

func (t *Transform) Use(key string, fn ResultFunc) {
	t.briefFnMap[key] = fn
}

func (t *Transform) Exec() {
	skipSeletors := len(t.selectorChan) == 0
	for _, card := range t.table.Cards {
		if !skipSeletors && !t.IsSelectCard(&card) {
			continue
		}
		for key, fn := range t.briefFnMap {
			t.result[key] = fn(t, t.result[key], &card)
		}
	}
}

func (t *Transform) IsSelectCard(c *models.Card) bool {
	for _, fn := range t.selectorChan {
		selectorFunc := fn(c)
		if !selectorFunc {
			return false
		}
	}
	return true
}

func (t *Transform) GetAllResult() map[string]interface{} {
	return t.result
}

func (t *Transform) GetResult(key string) interface{} {
	return t.result[key]
}

func (t *Transform) GetListById(id string) *models.List {
	value, ok := t.listsMap[id]
	if !ok {
		return nil
	}
	return value
}
