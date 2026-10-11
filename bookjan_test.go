package koa_test

import (
	"testing"

	"github.com/Sakaki-Aruka/koa"
)

func TestGetCcode(t *testing.T) {
	t.Run("Ccode position", func(t *testing.T) {
		// 2段目 192 0193 00600 1 -> Cコードは 0193
		got, err := koa.GetCcode("97840030007171920193006001")
		if err != nil || got != "0193" {
			t.Errorf("97840030007171920193006001 includes C0193 (%q, %v)", got, err)
		}
	})

	t.Run("Invalid content", func(t *testing.T) {
		// 2段目 191 0003 00000 8 -> チェックディジットは正しいが C0003 は欠番
		if _, err := koa.GetCcode("97840030007171910003000008"); err == nil {
			t.Error("97840030007171910003000008 includes an invalid content code (03)")
		}
	})
}

func TestGetPrice(t *testing.T) {
	t.Run("Leading zero", func(t *testing.T) {
		// 先頭の 0 を 8 進数として解釈しないこと (00600 -> 600)
		price, includeTax, err := koa.GetPrice("97840030007171920193006001")
		if err != nil || price != 600 || includeTax {
			t.Errorf("97840030007171920193006001 indicates 600 yen without tax (%d, %v, %v)", price, includeTax, err)
		}
	})

	t.Run("Leading zero with 8 and 9", func(t *testing.T) {
		// 2段目 192 0993 00980 6 -> 8 進数では読めない桁を含む
		price, _, err := koa.GetPrice("97840030007171920993009806")
		if err != nil || price != 980 {
			t.Errorf("97840030007171920993009806 indicates 980 yen (%d, %v)", price, err)
		}
	})

	t.Run("Tax included", func(t *testing.T) {
		// 2段目 191 0087 01000 9 -> 191 は税込
		price, includeTax, err := koa.GetPrice("97840030007171910087010009")
		if err != nil || price != 1000 || !includeTax {
			t.Errorf("97840030007171910087010009 indicates 1000 yen with tax (%d, %v, %v)", price, includeTax, err)
		}
	})
}
