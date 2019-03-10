package server

import "github.com/YuShuanHsieh/trello-transform/helpers/system"

type Configuration struct {
	FilePath string
	Port     int
}

type ConfigOption func(*Configuration)

var defaultConfiguration = Configuration{
	FilePath: "./",
	Port:     system.GetPortNumber(8181),
}

func Port(port int) ConfigOption {
	return func(config *Configuration) {
		config.Port = port
	}
}

func FilePath(path string) ConfigOption {
	return func(config *Configuration) {
		config.FilePath = path
	}
}
