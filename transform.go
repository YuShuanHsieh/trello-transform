package transform

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/adlio/trello"
	"go.uber.org/zap"
)

type SelectFunc func(*trello.Card) bool

type TransformFunc func(cards []*trello.Card) (key, result string, err error)

type results struct {
	m   map[string]string
	mux sync.Mutex
}

type Transform struct {
	client        *trello.Client
	trelloKey     string
	trelloToken   string
	trelloBoardID string

	board          trello.Board
	transformFuncs []TransformFunc
	selectFuncs    []SelectFunc
	result         results

	logger *zap.Logger
}

func New(logger *zap.Logger, key, token, id string) *Transform {
	return &Transform{
		client:        trello.NewClient(key, token),
		trelloKey:     key,
		trelloToken:   token,
		trelloBoardID: id,
		result:        results{m: make(map[string]string)},
		logger:        logger.With(zap.String("service", "transform")),
	}

}

func (t *Transform) AddSelect(s SelectFunc) {
	t.selectFuncs = append(t.selectFuncs, s)
}

func (t *Transform) AddTransformFunc(s TransformFunc) {
	t.transformFuncs = append(t.transformFuncs, s)
}

func (t *Transform) Exec() error {
	cards, err := t.getCards()
	if err != nil {
		return err
	}
	if len(t.selectFuncs) > 0 {
		cards = t.selectCards(cards)
	}

	var group sync.WaitGroup
	group.Add(len(t.transformFuncs))

	for _, fn := range t.transformFuncs {
		go func(cards []*trello.Card, transFn TransformFunc) {
			key, result, err := transFn(cards)
			if err != nil {
				t.logger.Error("failed to transform", zap.String("transform-func", key))
			}
			t.result.mux.Lock()
			t.result.m[key] = result
			t.result.mux.Unlock()
			group.Done()
		}(cards, fn)
	}

	group.Wait()
	t.printResult()

	return nil
}

func (t *Transform) printResult() {
	var builder strings.Builder
	for k, v := range t.result.m {
		builder.WriteString(fmt.Sprintf("[%s]\r\n%s\r\n", strings.ToUpper(k), v))
	}
	fmt.Println(builder.String())
}

func (t *Transform) selectCards(cards []*trello.Card) []*trello.Card {
	selectedCards := make([]*trello.Card, 0, len(cards)/2)

	for i, _ := range cards {
		var isSelected bool
		for _, fn := range t.selectFuncs {
			if fn(cards[i]) {
				isSelected = true
				break
			}
		}
		if isSelected {
			selectedCards = append(selectedCards, cards[i])
		}
	}

	t.logger.Debug("selected cards", zap.Int("card-number", len(selectedCards)))
	return selectedCards
}

func (t *Transform) getCards() ([]*trello.Card, error) {
	board, err := t.client.GetBoard(t.trelloBoardID, trello.Defaults())
	if err != nil {
		return nil, errors.Errorf("failed to get board %s: %v", t.trelloBoardID, err)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, errors.Errorf("failed to get cards: %v", err)
	}

	t.logger.Debug("fetched cards from trello api", zap.Int("card-number", len(cards)))

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, errors.Errorf("failed to get lists: %v", err)
	}

	listm := make(map[string]*trello.List)

	for i := range lists {
		listm[lists[i].ID] = lists[i]
	}

	for i, card := range cards {
		_, ok := listm[card.IDList]
		if ok {
			cards[i].List = listm[card.IDList]
		}
	}
	return cards, nil
}
