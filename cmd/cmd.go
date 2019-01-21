package cmd

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/services"
)

func Run() map[string]interface{} {
	filePath := flag.String("p", "trello.json", "The path of exported json file from Trello")
	if *filePath == "" {
		log.Panicln("The file path cannot be empty")
	}

	flag.Parse()

	content, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Panicf("Cannot open file: %s", *filePath)
	}

	tr := transform.New(content)
	tr.SelectorConfig(tr.SelectByList(&models.List{
		Name: "2019/01"}))
	tr.ResultConfig("list", transform.CardBriefFn)
	tr.ResultConfig("reference", transform.ExtractReferenceFn)
	tr.ResultConfig("label", transform.CountLabelsFn)
	tr.TransformFromTrello()

	return tr.GetAllResult()
}
