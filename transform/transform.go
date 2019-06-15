package transform

import (
	"encoding/json"
	"log"

	"github.com/YuShuanHsieh/trello-transform/logger"
	"github.com/YuShuanHsieh/trello-transform/trello"
)

// Seletor is function for filtering a card
type Seletor func(Context, *trello.Card) bool

// Accumulator is the collection of output value
type Accumulator interface{}

type Transformer func(Context, Accumulator, *trello.Card) (Accumulator, error)

type Transform struct {
	table        trello.Table
	ctx          Context
	transformers map[string]Transformer
	selectors    []Seletor
	result       map[string]interface{}

	logger logger.Logger
}

type Context struct {
	Labels map[string]trello.Label
	Lists  map[string]trello.List
}

func New(rawData []byte) *Transform {
	trans := Transform{
		ctx: Context{
			Lists:  make(map[string]trello.List),
			Labels: make(map[string]trello.Label),
		},
		transformers: make(map[string]Transformer),
		result:       make(map[string]interface{}),
		logger:       logger.GetLogger("Transform"),
	}

	err := json.Unmarshal(rawData, &trans.table)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	for _, v := range trans.table.Labels {
		trans.ctx.Labels[v.ID] = v
	}
	for _, v := range trans.table.Lists {
		trans.ctx.Lists[v.ID] = v
	}
	return &trans
}

func (t *Transform) Select(s Seletor) {
	if s != nil {
		t.selectors = append(t.selectors, s)
	}
}

func (t *Transform) Use(key string, fn Transformer) {
	t.transformers[key] = fn
}

func (t *Transform) Exec() {
	skipSeletors := len(t.selectors) == 0
	for _, card := range t.table.Cards {
		if !skipSeletors && !t.IsSelectCard(t.ctx, &card) {
			continue
		}
		for key, trans := range t.transformers {
			acc, err := trans(t.ctx, t.result[key], &card)
			if err != nil {
				t.logger.Error(err)
			}
			t.result[key] = acc
		}
	}
}

func (t *Transform) IsSelectCard(ctx Context, c *trello.Card) bool {
	for _, s := range t.selectors {
		if !s(ctx, c) {
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
