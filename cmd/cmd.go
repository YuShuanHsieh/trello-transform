package cmd

import (
	"flag"
	"io/ioutil"
	"log"
)

func Run() []byte {
	filePath := flag.String("p", "trello.json", "The path of exported json file from Trello")
	if *filePath == "" {
		log.Panicln("The file path cannot be empty")
	}

	flag.Parse()

	content, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Panicf("Cannot open file: %s", *filePath)
	}
	return content
}
