# Trello-Transform

## Description

A small project that transforms trello cards to plain text with target items.

## Why use it

When you have lots of cards on trello boards and just want to pick up parts of information from them.

## How to use

### Optionss

|  Option  |      Value      |  Description |
|----------|:-------------|------|
| key |  | the api key of trello([https://trello.com/app-key](https://trello.com/app-key)) |
| token |   | the develop token of trello(generated from [https://trello.com/app-key](https://trello.com/app-key)) |
| id |  | the target trello board id |
| list |  | the name of lists which would be parsed |
| target | `titles`,`links` | the target types of output |

CLI Example:
`go run cmd.go --key <your key> --token <your token> --id <board id> --list 2020/05 --target titles --target links`
