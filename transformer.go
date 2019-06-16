package transform

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/YuShuanHsieh/trello-transform/trello"
)

// CardHeader is the header of a card
type CardHeader struct {
	Date  string `json:"date"`
	Title string `json:"title"`
}

type markdownURLs []string

func (m markdownURLs) String() string {
	var builder strings.Builder
	for _, v := range m {
		builder.WriteString("- " + v + "\r\n")
	}
	return builder.String()
}

// MarkdownURLTransformer extract url with markdowm style
// e.g. [content](www.example.com)
func MarkdownURLTransformer(
	ctx Context,
	acc Accumulator,
	c *trello.Card,
) (Accumulator, error) {
	value, ok := acc.(markdownURLs)
	if !ok {
		value = markdownURLs{}
	}
	reg, err := regexp.Compile(`[[][^]]+[]][(][^)]+[)]`)
	if err != nil {
		return acc, err
	}
	targets := reg.FindAllString(c.Desc, -1)
	return append(value, targets...), nil
}

type countlabel map[string]int

func (c countlabel) String() string {
	var builder strings.Builder
	for k, v := range c {
		builder.WriteString(k + ":" + strconv.Itoa(v) + "\r\n")
	}
	return builder.String()
}

// CountLabelsTransformer Count label numbers for all cards.
// Each card may has more than one label.
func CountLabelsTransformer(
	ctx Context,
	acc Accumulator,
	c *trello.Card,
) (Accumulator, error) {
	value, ok := acc.(countlabel)
	if !ok {
		value = make(countlabel)
	}
	for _, id := range c.IDLabels {
		v, ok := ctx.Labels[id]
		if ok {
			value[v.Name]++
		}
	}
	return value, nil
}

type headers []CardHeader

func (h headers) String() string {
	var builder strings.Builder
	for _, v := range h {
		builder.WriteString("- " + v.Date + " " + v.Title + "\r\n")
	}
	return builder.String()
}

// CardHeaderTransformer Generate a short sentence with card's title and date.
func CardHeaderTransformer(
	ctx Context,
	acc Accumulator,
	c *trello.Card,
) (Accumulator, error) {
	// Ignore if the card has no data or title
	if isEmptyString(c.Due) || isEmptyString(c.Name) {
		return acc, nil
	}
	v, ok := acc.(headers)
	if !ok {
		v = headers{}
	}
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
