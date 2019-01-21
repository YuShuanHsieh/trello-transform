package main

import (
	"log"

	"github.com/YuShuanHsieh/trello-transform/cmd"
)

func main() {
	result := cmd.Run()
	log.Printf("%+v", result)
}
