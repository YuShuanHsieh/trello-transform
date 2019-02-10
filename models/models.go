package models

type SeletorFunc func(*Card) bool

type TrelloTable struct {
	Name   string  `json:"name"`
	Cards  []Card  `json:"cards"`
	Labels []Label `json:"labels"`
	Lists  []List  `json:"lists"`
}

type List struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}

type Card struct {
	IDLabels []string `json:"idLabels"`
	IDList   string   `json:"idList"`
	Name     string   `json:"name"`
	Due      string   `json:"due"`
	Desc     string   `json:"desc"`
	Closed   bool     `json:"closed"`
}

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
