package koa_test

import (
	"fmt"
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
		if err != nil || category == "" || content == "" {
			t.Error("C0000 includes a valid content code (00)")
		}
	})
}

// 日本図書コード管理センター「分類記号・各コードの内容解説・注意事項」の内容コード (3・4桁目)
var officialContents = map[string]string{
	"00": "総記", "01": "百科事典", "02": "年鑑・雑誌", "04": "情報科学",
	"10": "哲学", "11": "心理 (学)", "12": "倫理 (学)", "14": "宗教", "15": "仏教", "16": "キリスト教",
	"20": "歴史総記", "21": "日本歴史", "22": "外国歴史", "23": "伝記", "25": "地理", "26": "旅行",
	"30": "社会科学総記", "31": "政治 - 含む国防軍事", "32": "法律", "33": "経済・財政・統計", "34": "経営", "36": "社会", "37": "教育", "39": "民族・風習",
	"40": "自然科学総記", "41": "数学", "42": "物理学", "43": "化学", "44": "天文・地学", "45": "生物学", "47": "医学・歯学・薬学",
	"50": "工学・工学総記", "51": "土木", "52": "建築", "53": "機械", "54": "電気", "55": "電子通信", "56": "海事", "57": "採鉱・冶金", "58": "その他の工業",
	"60": "産業総記", "61": "農林業", "62": "水産業", "63": "商業", "65": "交通・通信",
	"70": "芸術総記", "71": "絵画・彫刻", "72": "写真・工芸", "73": "音楽・舞踊", "74": "演劇・映画", "75": "体育・スポーツ", "76": "諸芸・娯楽", "77": "家事", "78": "日記・手帳", "79": "コミックス・劇画",
	"80": "語学総記", "81": "日本語", "82": "英米語", "84": "ドイツ語", "85": "フランス語", "87": "各国語",
	"90": "文学総記", "91": "日本文学総記", "92": "日本文学詩歌", "93": "日本文学、小説・物語", "95": "日本文学、評論、随筆、その他", "97": "外国文学小説", "98": "外国文学、その他",
}

func TestCContentTable(t *testing.T) {
	for i := 0; i < 100; i++ {
		code := fmt.Sprintf("%02d", i)
		want, ok := officialContents[code]
		_, got, err := koa.CContent("C00" + code)
		if ok && (err != nil || got != want) {
			t.Errorf("C00%s should be %q (%q, %v)", code, want, got, err)
		}
		if !ok && err == nil {
			t.Errorf("C00%s should be invalid (%q)", code, got)
		}
	}
}
