package transform

import (
	"testing"

	"github.com/YuShuanHsieh/trello-transform/trello"
)

func TestCompareList(t *testing.T) {
	target := trello.List{
		Closed: true,
		Name:   "TestList",
	}

	results := []struct {
		list     trello.List
		expected bool
	}{
		{
			trello.List{
				Name: "Book",
			},
			false,
		},
		{
			trello.List{
				Closed: true,
				Name:   "Test",
			},
			true,
		},
		{
			trello.List{
				Closed: true,
				Name:   "TestList",
			},
			true,
		},
		{
			trello.List{
				Closed: false,
				Name:   "TestList",
			},
			false,
		},
	}

	for _, result := range results {
		test := compareList(target, result.list)
		if test != result.expected {
			t.Errorf("Expected %t but got %t \n", result.expected, test)
		}
	}
}

func TestSelectByNames(t *testing.T) {
	ctx := Context{
		Lists: map[string]trello.List{"id-1": trello.List{Name: "List1"}},
	}
	tests := []struct {
		selector Seletor
		card     trello.Card
		expect   bool
	}{
		{
			selector: ByListNames("List1", "List2", "List3"),
			card: trello.Card{
				IDList: "id-1",
			},
			expect: true,
		},
		{
			selector: ByListNames("List2", "List3"),
			card: trello.Card{
				IDList: "id-1",
			},
			expect: false,
		},
	}
	var result bool
	for _, test := range tests {
		result = test.selector(ctx, &test.card)
		if result != test.expect {
			t.Errorf("Expected %t but got %t", test.expect, result)
		}
	}
}
