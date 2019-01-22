package main

import (
	"github.com/YuShuanHsieh/trello-transform/server"
)

func main() {
	s := server.Default()
	s.Run()
}
