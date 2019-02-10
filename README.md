# Trello-Transform

## Description

Simple project for transform json file exported from trello to customized output.

## Why use it

1. Sometime You just need few data from trello cards. Use Trello API to retrieve data is too heavy work to you.
2. When you create lots of cards on trello board and want to group parts of information from them.

## How to use / test

- export json file from your card (for free users)
- run `go run main.go` to start the web server
- run `go run cmd/cmd.go <your trello json file> | json_pp`

### How to change to code

```go
// content is []byte type from json file
tr := transform.New(content)

// Now selector function only support `SelectByList`
tr.Selector(tr.SelectByList(&models.List{Name: "2019/01"}))

// `CardBriefFunc`, `ExtractReferenceFunc`, `CountLabelsFunc` are default fn for service. Or you can create your result config function
tr.Use("list", transform.CardBriefFunc)
tr.Use("reference", transform.ExtractReferenceFunc)
tr.Use("label", transform.CountLabelsFunc)

// Call this function to transform cards
tr.Exec()

json, err := json.Marshal(tr.GetAllResult())
if err != nil {
	log.Printf(err.Error())
}
log.Printf("%s", json)
```

## Create a customized `ResultConfig` function

```go
type ResultConfigFn func(*Transform, interface{}, *models.Card) interface{}

func CountLabelsFn(tr *Transform, preValue interface{}, c *models.Card) interface{} {
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
```

## Config selector function

```go
// Cards with target list name will be transformed.
// Add this line before call `tr.TransformFromTrello()`
tr.Selector(tr.SelectByList(&models.List{Name: "2019/01"}))
``` 
