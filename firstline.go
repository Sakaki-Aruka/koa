package koa

import "regexp"

var firstLinePattern *regexp.Regexp = regexp.MustCompile(`^978([0-9]{9})([0-9])$`)

func FIsValid(s string) bool {
	if !firstLinePattern.MatchString(s) {
		return false
	}
	actual := checkdigit(s)
	d, _ := byteToInt(s[12])
	return actual == d
}
