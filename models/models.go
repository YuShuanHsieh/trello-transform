package models

type TrelloTable struct {
	Name   string  `json:"name"`
	Cards  []Card  `json:"cards"`
	Labels []Label `json:"labels"`
}

type Card struct {
	IDLabels []string `json:"idLabels"`
	IDList   string   `json:"idList"`
	Name     string   `json:"name"`
	Due      string   `json:"due"`
	Desc     string   `json:"desc"`
}

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
