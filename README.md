# Trello-Transform

## Description

A small project that transforms json file exported from trello to customized content.

## Why use it

- Sometime You just need few data from trello cards. Use Trello API to retrieve data is too heavy work to you.
- When you have lots of cards on trello boards and just want to pick up parts of information from them.

## How to use

- Export json files from your Trello table (this feature is available for all users)
- run `go run cmd/cmd.go -file=<local file path> -types=<transformer type> --lists=<list names>`

### Optionss

|  Option  |      Type      |  Description |
|----------|:-------------|------|
| file(required) |  `string` | A local path of json file |
| lists |    `string`   |   Select cards with matched list names. format: `Name1,Name2`. Separate each list name by `,`. |
| types |  `string` | Supported types: `headers`, `labels`, `urls`. Separate each type by `,`. |

CLI Example:
`go run cmd/cmd.go -file="./example.json" -lists="2019/06" -types="urls,headers"`

## Create a `Transformer`

```go
// Transformer Type definition
type Transformer func(Context, Accumulator, *trello.Card) (Accumulator, error)

// Implement Accumulator interface
type example string
func (e example) String() string {
	return e
}

// Implement your Transformer
func Transformer(transform.Context,acc transform.Accumulator,c *trello.Card) (transform.Accumulator, error) {
	value, ok := acc.(example)
	// To prevent get nil value, we need to check the type asset is correct.
	if !ok {
		value = example{}
	}
	// do somthing to modify value...

	return value, nil
}
```