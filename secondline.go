package koa

import (
	"fmt"
	"regexp"
	"strconv"
)

var invalidSecondLineFormat error = fmt.Errorf("Invalid book jan second line format")

var secondLinePattern *regexp.Regexp = regexp.MustCompile(`^19(1|2)([0-9]{4})([0-9]{5})([0-9])$`)

func SIsValid(s string) bool {
	if !secondLinePattern.MatchString(s) {
		return false
	}
	actual := checkdigit(s)
	d := int(s[12] - '0')
	return actual == d
}

// SGetPrice, includeTax, error
func SGetPrice(s string) (uint32, bool, error) {
	if !SIsValid(s) {
		return 0, false, invalidSecondLineFormat
	}

	match := secondLinePattern.FindStringSubmatch(s)
	includeTax := match[1] == "1"
	var price uint32
	p, err := strconv.ParseUint(match[3], 10, 32)
	if err != nil {
		return 0, false, err
	}
	price = uint32(p)
	return price, includeTax, nil
}

func SGetCcode(s string) (string, error) {
	if !SIsValid(s) {
		return "", invalidSecondLineFormat
	}
	match := secondLinePattern.FindStringSubmatch(s)
	c := match[2]
	if CIsValid(c) {
		return c, nil
	} else {
		return "", invalidCcodeFormat
	}
}
