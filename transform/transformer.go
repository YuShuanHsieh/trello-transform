package transform

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/YuShuanHsieh/trello-transform/trello"
)

// CardHeader is the header of a card
type CardHeader struct {
	Date  string `json:"date"`
	Title string `json:"title"`
}

// MarkdownURLTransformer extract url with markdowm style
// e.g. [content](www.example.com)
func MarkdownURLTransformer(
	ctx *Context,
	acc Accumulator,
	c *trello.Card,
) (Accumulator, error) {
	value := acc.([]string)
	reg, err := regexp.Compile(`[[][\S\s]+[]][(][\S]+[)]`)
	if err != nil {
		return acc, err
	}
	targets := reg.FindAllString(c.Desc, -1)
	return append(value, targets...), nil
}

// CountLabelsTransformer Count label numbers for all cards.
// Each card may has more than one label.
func CountLabelsTransformer(
	ctx *Context,
	acc Accumulator,
	c *trello.Card,
) (Accumulator, error) {
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

// CardHeaderTransformer Generate a short sentence with card's title and date.
func CardHeaderTransformer(
	ctx *Context,
	acc Accumulator,
	c *trello.Card,
) (Accumulator, error) {
	// Ignore if the card has no data or title
	if isEmptyString(c.Due) || isEmptyString(c.Name) {
		return acc, nil
	}
	v := acc.([]CardHeader)
	t, err := time.Parse(time.RFC3339, c.Due)
	if err != nil {
		return acc, err
	}

	return append(v, CardHeader{
		Date:  fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day()),
		Title: c.Name,
	}), nil
}

func isEmptyString(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}
