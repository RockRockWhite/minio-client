package utils

import (
	"regexp"
)

var reg *regexp.Regexp

func init() {
	reg, _ = regexp.Compile("(.*)\\.(.*)")
}

func GetPrefixAndPosfix(objectname string) (string, string) {
	submatch := reg.FindStringSubmatch(objectname)
	if len(submatch) != 3 {
		return "", ""
	}
	return submatch[1], submatch[2]
}
