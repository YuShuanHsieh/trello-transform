package main

import (
	"encoding/json"
	"log"

	"github.com/YuShuanHsieh/trello-transform/cmd"
	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/services"
)

func main() {
	content := cmd.Run()

	tr := transform.New(content)
	tr.SelectorConfig(tr.SelectByList(&models.List{
		Name: "2019/01"}))
	tr.ResultConfig("list", transform.CardBriefFn)
	tr.ResultConfig("reference", transform.ExtractReferenceFn)
	tr.ResultConfig("label", transform.CountLabelsFn)
	tr.TransformFromTrello()

	json, err := json.Marshal(tr.GetAllResult())
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("%s", json)
}
