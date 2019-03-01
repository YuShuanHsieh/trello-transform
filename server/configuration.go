package server

import "github.com/YuShuanHsieh/trello-transform/system"

const (
	DefaultFilePath = "./"
	DefaultPort     = 8181
)

type ServerConfiguration struct {
	FilePath string
	Port     int
}

func defaultConfiguration() *ServerConfiguration {
	p := system.GetPortNumber(DefaultPort)
	return &ServerConfiguration{
		FilePath: DefaultFilePath,
		Port:     p}
}
