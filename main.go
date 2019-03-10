package main

import (
	"github.com/YuShuanHsieh/trello-transform/server"
)

func main() {
	srv := server.New(server.Port(8080))
	srv.Run()
}
