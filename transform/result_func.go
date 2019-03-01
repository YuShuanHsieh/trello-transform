package transform

import (
	"fmt"
	"regexp"
	"time"

	"github.com/YuShuanHsieh/trello-transform/errors"
	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/validators"
)

// ResultFunc generate a result value from each card
type ResultFunc func(*Transform, interface{}, *models.Card) interface{}

// ExtractReferenceFunc extract references with markdown style
func ExtractReferenceFunc(tr *Transform, preValue interface{}, c *models.Card) interface{} {
	arr, ok := preValue.([]string)
	if !ok {
		arr = []string{}
	}
	reg, err := regexp.Compile(`[[][\S\s]+[]][(][\S]+[)]`)
	if err != nil {
		errors.Log(err.Error())
		return nil
	}
	targets := reg.FindAllString(c.Desc, -1)
	arr = append(arr, targets...)

	return arr
}

// CountLabelsFunc Count label numbers. Each card has more than a label.
func CountLabelsFunc(tr *Transform, preValue interface{}, c *models.Card) interface{} {
	labelsMap, ok := preValue.(map[string]int)
	if !ok {
		labelsMap = make(map[string]int)
	}
	for _, id := range c.IDLabels {
		v, ok := tr.labelsMap[id]
		if ok {
			labelsMap[v.Name]++
		}
	}

	return labelsMap
}

// CardBriefFunc Generate a short sentence from a card. This is a example to combine date with title.
func CardBriefFunc(tr *Transform, preValue interface{}, c *models.Card) interface{} {
	if validators.IsEmptyString(c.Due) || validators.IsEmptyString(c.Name) {
		return nil
	}

	v, ok := preValue.([]string)

	if !ok {
		v = []string{}
	}

	t, err := time.Parse(time.RFC3339, c.Due)
	if err != nil {
		errors.Log(err.Error())
		return nil
	}

	item := fmt.Sprintf("%d %d %d - %s", t.Year(), t.Month(), t.Day(), c.Name)

	v = append(v, item)
	return v
}
