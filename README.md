# Trello-Transform

## Description

Small project for transform json file exported from trello to customized output.

## Why to use
When you create lots of cards on trello board and want to group important information from them.

## How to use

`go run main.go -p ./trello.json`

- `-p`: the path of your json file

## How it work

1. Export `json file` from trello board
2. Use `json file` as input
3. Define your ResultConfigs
4. Run transform function
5. Call the result function get get all or single result

### Example

```go
// content is []byte type from json file
tr := transform.New(content)

// Now selector function only support `SelectByList`
tr.SelectorConfig(tr.SelectByList(&models.List{Name: "2019/01"}))

// `CardBriefFn`, `ExtractReferenceFn`, `CountLabelsFn` are default fn for service. Or you can create your result config function
tr.ResultConfig("list", transform.CardBriefFn)
tr.ResultConfig("reference", transform.ExtractReferenceFn)
tr.ResultConfig("label", transform.CountLabelsFn)

// Call this function to transform cards
tr.TransformFromTrello()

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
tr.SelectorConfig(tr.SelectByList(&models.List{Name: "2019/01"}))
``` 
