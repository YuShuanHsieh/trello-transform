package utilities

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsEmptyString(value string) bool {
	return strings.TrimSpace(value) == ""
}

func IsFromLocalHost(remoteAddr string) bool {
	matched, _ := regexp.MatchString("^\\[::1\\]", remoteAddr)
	return matched
}

func GetPortNumber(defaultPort int) int {
	var p int
	port, isExisted := os.LookupEnv("SERVER_PORT")

	if p = defaultPort; isExisted {
		p, _ = strconv.Atoi(port)
	}

	return p
}
