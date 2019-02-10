package utilities

import (
	"os"
	"strconv"
	"strings"
)

func GetPortNumber(defaultPort int) int {
	var p int
	port, isExisted := os.LookupEnv("SERVER_PORT")

	if p = defaultPort; isExisted {
		p, _ = strconv.Atoi(port)
	}

	return p
}

func IsEmptyString(value string) bool {
	return strings.TrimSpace(value) == ""
}
