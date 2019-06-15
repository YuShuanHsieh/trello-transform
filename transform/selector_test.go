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
	var ctx Context
	tests := []struct {
		selector Seletor
		card     trello.Card
		expect   bool
	}{
		{
			selector: ByListNames("Card1", "Card2", "Card3"),
			card: trello.Card{
				Name: "Card1",
			},
			expect: true,
		},
		{
			selector: ByListNames("Card2", "Card3"),
			card: trello.Card{
				Name: "Card1",
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
