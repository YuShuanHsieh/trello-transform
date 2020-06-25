package transform

import (
	"regexp"
	"strings"

	"github.com/adlio/trello"
)

func ToTitles(cards []*trello.Card) (string, string, error) {
	key := "titles"
	var builder strings.Builder
	for _, card := range cards {
		if card.Name != "" {
			builder.WriteString("-" + card.Due.Format("2006/01/02") + " " + card.Name + "\r\n")
		}
	}
	return key, builder.String(), nil
}

func ToLinks(cards []*trello.Card) (string, string, error) {
	key := "links"
	var builder strings.Builder
	reg, _ := regexp.Compile(`[[][^]]+[]][(][^)]+[)]`)
	for _, card := range cards {
		if card.Desc != "" {
			links := reg.FindAllString(card.Desc, -1)
			for _, link := range links {
				builder.WriteString("-" + link + "\r\n")
			}
		}
	}
	return key, builder.String(), nil
}
