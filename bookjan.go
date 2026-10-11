package koa

import (
	"errors"
)

var invalidBookJan error = errors.New("Invalid Book JAN code")

func checkdigit(s string) int {
	sum := 0
	for i := 0; i < 12; i++ {
		n, err := byteToInt(s[i])
		if err != nil {
			panic(err.Error())
		}
		if i&1 == 0 {
			sum += n
		} else {
			sum += n * 3
		}
	}
	return (10 - (sum % 10)) % 10
}

func byteToInt(b byte) (int, error) {
	if b < '0' || '9' < b {
		return -1, errors.New("Invalid character")
	}
	return int(b - '0'), nil
}

func IsValidBookJan(s string) bool {
	if len(s) != 26 {
		return false
	}
	return FIsValid(s[:13]) && SIsValid(s[13:])
}

func GetCcode(s string) (string, error) {
	if !IsValidBookJan(s) {
		return "", invalidBookJan
	}
	return SGetCcode(s[13:])
}

// Price (yen), Tax included, error
func GetPrice(s string) (uint32, bool, error) {
	if !IsValidBookJan(s) {
		return 0, false, invalidBookJan
	}
	p, t, err := SGetPrice(s[13:])
	if err != nil {
		return 0, false, err
	}
	return p, t, nil
}
