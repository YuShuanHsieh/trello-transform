# Trello-Transform

## Description

A simple project for transform json file exported from trello to customized output.

## Why use it

1. Sometime You just need few data from trello cards. Use Trello API to retrieve data is too heavy work to you.
2. When you create lots of cards on trello board and want to pick up parts of information from them.

## How to use / test

- Export json files from your Trello table (this feature is for free users)
- run `go run main.go` to start the web server
- run `go run cmd/cmd.go <your trello json file> | json_pp` (Website is preparing)

### How to add more handler

```go
// The type of content is []byte
tr := transform.New(content)

// Add a selector for filter cards.
tr.Select(selector.ByList(trello.List{name: "2019/02"}))

// Use handlers to pick up information from each card
tr.Use("list", defaultHandler.CardBriefHandler)
tr.Use("reference", defaultHandler.ExtractReferenceHandler)
tr.Use("label", defaultHandler.CountLabelsHandler)

// Call this function to transform cards
tr.Exec()
```

## Create a `TransformHandler`

```go
type TransformHandler func(*Context, Accumulator, *trello.Card) (Accumulator, error)

func ExtractReferenceHandler(
	ctx *transform.Context,
	acc transform.Accumulator,
	c *trello.Card,
) (transform.Accumulator, error) {
	// Check the underlying data type first.
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
```

## Config selector function

```go
// Cards with target list name will be transformed.
// Add this line before call `tr.Exec()`
tr.Select(selector.ByList(trello.List{name: "2019/02"}))
``` 
