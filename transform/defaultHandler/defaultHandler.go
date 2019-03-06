package defaultHandler

import (
	"fmt"
	"regexp"
	"time"

	"github.com/YuShuanHsieh/trello-transform/errors"
	"github.com/YuShuanHsieh/trello-transform/helpers/validators"
	"github.com/YuShuanHsieh/trello-transform/transform"
	"github.com/YuShuanHsieh/trello-transform/trello"
)

type CardBrief struct {
	Date  string `json:"date"`
	Title string `json:"title"`
}

// ExtractReferenceHandler extract references with markdown style
func ExtractReferenceHandler(
	ctx *transform.Context,
	acc transform.Accumulator,
	c *trello.Card,
) (transform.Accumulator, error) {
	value, ok := acc.([]string)
	if !ok {
		value = []string{}
	}
	reg, err := regexp.Compile(`[[][\S\s]+[]][(][\S]+[)]`)
	if err != nil {
		return acc, err
	}
	targets := reg.FindAllString(c.Desc, -1)
	return append(value, targets...), nil
}

// CountLabelsHandler Count label numbers. Each card has more than a label.
func CountLabelsHandler(
	ctx *transform.Context,
	acc transform.Accumulator,
	c *trello.Card,
) (transform.Accumulator, error) {
	value, ok := acc.(map[string]int)
	if !ok {
		value = make(map[string]int)
	}
	for _, id := range c.IDLabels {
		v, ok := ctx.Labels[id]
		if ok {
			value[v.Name]++
		}
	}
	return value, nil
}

// CardBriefHandler Generate a short sentence from a card. This is a example to combine date with title.
func CardBriefHandler(
	ctx *transform.Context,
	acc transform.Accumulator,
	c *trello.Card,
) (transform.Accumulator, error) {
	if validators.IsEmptyString(c.Due) || validators.IsEmptyString(c.Name) {
		return acc, nil
	}

	v, ok := acc.([]CardBrief)
	if !ok {
		v = []CardBrief{}
	}

	t, err := time.Parse(time.RFC3339, c.Due)
	if err != nil {
		errors.Log(err.Error())
		return acc, err
	}

	return append(v, CardBrief{
		Date:  fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day()),
		Title: c.Name,
	}), nil
}
