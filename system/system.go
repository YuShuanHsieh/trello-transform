package system

import (
	"os"
	"strconv"
)

func GetPortNumber(defaultPort int) int {
	var p int
	port, isExisted := os.LookupEnv("SERVER_PORT")

	if p = defaultPort; isExisted {
		p, _ = strconv.Atoi(port)
	}

	return p
}
