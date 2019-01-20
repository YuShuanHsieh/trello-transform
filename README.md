# Trello-Transform

## Description

Small project for transform json file exported from trello to customized output.

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
tr.ResultConfig("list", transform.CardBriefFn)
tr.ResultConfig("reference", transform.ExtractReferenceFn)
tr.ResultConfig("label", transform.CountLabelsFn)
tr.TransformFromTrello()

log.Printf("%+v", tr.GetAllResult())
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
