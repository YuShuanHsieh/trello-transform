package server

const (
	defaultFilePath = "./"
	defaultPort     = 8181
)

type ServerConfiguration struct {
	FilePath string
	Port     int
}

func defaultConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		FilePath: defaultFilePath,
		Port:     defaultPort}
}
