package koa_test

import (
	"testing"

	"github.com/Sakaki-Aruka/koa"
)

func TestSIsValid(t *testing.T) {
	t.Run("Invalid checkdigit", func(t *testing.T) {
		got := koa.SIsValid("1920193006000")
		if got {
			t.Error("1920193006000 has an invalid checkdigit")
		}
	})

	t.Run("Valid checkdigit", func(t *testing.T) {
		got := koa.SIsValid("1920193006001")
		if !got {
			t.Error("1920193006001 has a valid checkdigit")
		}
	})
}