package transform

import (
	"encoding/json"
	"log"

	"github.com/YuShuanHsieh/trello-transform/trello"
)

type Seletor func(*Context, *trello.Card) bool

type Accumulator interface{}

type TransformHandler func(*Context, Accumulator, *trello.Card) (Accumulator, error)

type Transform struct {
	table     trello.Table
	labels    map[string]*trello.Label
	lists     map[string]*trello.List
	handlers  map[string]TransformHandler
	selectors []Seletor
	result    map[string]interface{}
}

type Context struct {
	Labels map[string]trello.Label
	Lists  map[string]trello.List
}

func New(rawData []byte) *Transform {
	t := initTransform()

	err := UnmarshalJson(rawData, &t.table)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	t.allocateLabels()
	t.allocateLists()
	return t
}

func initTransform() *Transform {
	return &Transform{
		lists:    make(map[string]*trello.List),
		labels:   make(map[string]*trello.Label),
		handlers: make(map[string]TransformHandler),
		result:   make(map[string]interface{})}
}

func (t *Transform) allocateLabels() {
	for i, v := range t.table.Labels {
		t.labels[v.ID] = &t.table.Labels[i]
	}
}

func (t *Transform) allocateLists() {
	for i, v := range t.table.Lists {
		t.lists[v.ID] = &t.table.Lists[i]
	}
}

func (t *Transform) Select(s Seletor) {
	if s != nil {
		t.selectors = append(t.selectors, s)
	}
}

func (t *Transform) Use(key string, fn TransformHandler) {
	t.handlers[key] = fn
}

func (t *Transform) Exec() {
	ctx := t.createContext()
	skipSeletors := len(t.selectors) == 0
	for _, card := range t.table.Cards {
		if !skipSeletors && !t.IsSelectCard(ctx, &card) {
			continue
		}
		for key, handler := range t.handlers {
			acc, err := handler(ctx, t.result[key], &card)
			if err != nil {
				// log errors
			}
			t.result[key] = acc
		}
	}
}

func (t *Transform) createContext() *Context {
	labels := make(map[string]trello.Label)
	lists := make(map[string]trello.List)
	for key, value := range t.labels {
		labels[key] = *value
	}
	for key, value := range t.lists {
		lists[key] = *value
	}
	return &Context{
		Lists:  lists,
		Labels: labels,
	}
}

func (t *Transform) IsSelectCard(ctx *Context, c *trello.Card) bool {
	for _, s := range t.selectors {
		selectorFunc := s(ctx, c)
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

func UnmarshalJson(jsonData []byte, t *trello.Table) error {
	err := json.Unmarshal(jsonData, t)
	if err != nil {
		return err
	}
	return nil
}
