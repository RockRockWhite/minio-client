package utils

import (
	"regexp"
)

var reg *regexp.Regexp

func init() {
	reg, _ = regexp.Compile("(.*)\\.(.*)")
}

func GetFileNameAndPosfix(filename string) (string, string) {
	submatch := reg.FindStringSubmatch(filename)
	if len(submatch) != 3 {
		return "", ""
	}
	return submatch[1], submatch[2]
}
