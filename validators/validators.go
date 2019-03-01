package validators

import (
	"regexp"
	"strings"
)

func IsEmptyString(value string) bool {
	return strings.TrimSpace(value) == ""
}

func IsFromLocalHost(remoteAddr string) bool {
	matched, _ := regexp.MatchString("^\\[::1\\]", remoteAddr)
	return matched
}
