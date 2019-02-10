package selector

import (
	"github.com/YuShuanHsieh/trello-transform/models"
	"testing"
)

func TestCompareList(t *testing.T) {
	target := models.List{
		Closed: true,
		Name:   "TestList",
	}

	results := []struct {
		list     models.List
		expected bool
	}{
		{
			models.List{
				Name: "Book",
			},
			false,
		},
		{
			models.List{
				Closed: true,
				Name:   "Test",
			},
			true,
		},
		{
			models.List{
				Closed: true,
				Name:   "TestList",
			},
			true,
		},
		{
			models.List{
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
