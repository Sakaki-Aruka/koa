package koa_test

import (
	"testing"

	"github.com/Sakaki-Aruka/koa"
)

func TestCIsValid(t *testing.T) {
	t.Run("Invalid prefix", func(t *testing.T) {
		got := koa.CIsValid("B0000")
		if got {
			t.Error("B0000 is invalid")
		}
	})

	t.Run("Invalid short", func(t *testing.T) {
		got := koa.CIsValid("C000")
		if got {
			t.Error("C000 is invalid")
		}
	})

	t.Run("Valid C prefix", func(t *testing.T) {
		got := koa.CIsValid("C0000")
		if !got {
			t.Error("C0000 is valid")
		}
	})

	t.Run("Valid c prefix", func(t *testing.T) {
		got := koa.CIsValid("c0000")
		if !got {
			t.Error("c0000 is valid")
		}
	})

	t.Run("Get target", func(t *testing.T) {
		_, err := koa.CTargetAudience("C0000")
		if err != nil {
			t.Error("C0000 includes a valid target code (0)")
		}
	})

	t.Run("Get form", func(t *testing.T) {
		_, err := koa.CForm("C0000")
		if err != nil {
			t.Error("C0000 includes a valid form code (0)")
		}
	})

	t.Run("Invalid content", func(t *testing.T) {
		_, _, err := koa.CContent("C0003")
		if err == nil {
			t.Error("C0003 includes an invalid content code (03)")
		}
	})

	t.Run("Get content", func(t *testing.T) {
		category, content, err := koa.CContent("C0000")
		if err != nil && category != "" && content != "" {
			t.Error("C0000 includes a valid content code (00)")
		}
	})
}