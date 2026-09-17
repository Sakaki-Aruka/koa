package koa_test

import (
	"testing"

	"github.com/Sakaki-Aruka/koa"
)

func TestFIsValid(t *testing.T) {
	t.Run("Invalid Japanese ISBN", func(t *testing.T) {
		got := koa.FIsValid("9794003000717")
		if got {
			t.Error("9794003000717 is not a Japanese ISBN")
		}
	})

	t.Run("Valid Japanese ISBN", func(t *testing.T) {
		got := koa.FIsValid("9784003000717")
		if !got {
			t.Error("9784003000717 is a valid Japanese ISBN")
		}
	})
}