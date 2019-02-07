package server

import "github.com/YuShuanHsieh/trello-transform/utilities"

const (
	DefaultFilePath = "./"
	DefaultPort     = 8181
)

type ServerConfiguration struct {
	FilePath string
	Port     int
}

func defaultConfiguration() *ServerConfiguration {
	p := utilities.GetPortNumber(DefaultPort)
	return &ServerConfiguration{
		FilePath: DefaultFilePath,
		Port:     p}
}
