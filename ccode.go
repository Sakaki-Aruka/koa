package koa

import (
	"fmt"
	"regexp"
	"strings"
)

var invalidCcodeFormat error = fmt.Errorf("Invalid code (0000~9998)")

var ccodePattern *regexp.Regexp = regexp.MustCompile(`^[Cc]?([0-9])([0-9])([0-9]{2})$`)

var targets []string = []string{"一般", "教養", "実用", "専門", "検定教科書・消費税非課税品・その他", "女性", "学参1 (小中)", "学参2 (高校)", "児童", "雑誌扱い"}
var forms []string = []string{"単行本", "文庫", "新書", "全集・双書", "ムック・その他", "事・辞典", "図鑑", "絵本", "磁性媒体など", "コミック"}
var contents map[byte][]string = map[byte][]string{
	// index=0 は種別名
	'0': {"総記", "総記", "百科事典", "年鑑・雑誌", "", "情報科学"},
	'1': {"哲学・宗教・心理学", "哲学", "心理 (学)", "倫理 (学)", "", "宗教", "仏教", "キリスト教"},
	'2': {"歴史・地理", "歴史総記", "日本歴史", "外国歴史", "伝記", "", "地理", "旅行"},
	'3': {"社会科学", "社会科学総記", "政治 - 含む国防軍事", "法律", "経済・財政・統計", "経営", "", "社会", "教育", "", "民族・風習"},
	'4': {"自然科学", "自然科学総記", "数学", "物理学", "化学", "天文・地学", "生物学", "", "医学・歯学・薬学"},
	'5': {"工学工業", "工学・工学総記", "土木", "建築", "機械", "電気", "電子通信", "海事", "採鉱・冶金", "その他の工業"},
	'6': {"産業", "産業総記", "農林業", "水産業", "商業", "", "交通・通信"},
	'7': {"芸術・生活", "芸術総記", "絵画・彫刻", "写真・工芸", "音楽・舞踊", "演劇・映画", "体育・スポーツ", "諸芸・娯楽", "家事", "日記・手帳", "コミックス・劇画"},
	'8': {"語学", "語学総記", "日本語", "英米語", "", "ドイツ語", "フランス語", "", "各国語"},
	'9': {"文学", "文学総記", "日本文学総記", "日本文学詩歌", "日本文学、小説・物語", "", "日本文学、評論、随筆、その他", "", "外国文学小説", "外国文学、その他"},
}

func roughCheck(c string) bool {
	return ccodePattern.MatchString(c)
}

var creplacer strings.Replacer = *strings.NewReplacer("C", "", "c", "")
func pureCcode(c string) string {
	return creplacer.Replace(c)
}

func CIsValid(c string) bool {
	if !roughCheck(c) {
		return false
	}

	if _, err := CTargetAudience(c); err != nil {
		return false
	}

	if _, err := CForm(c); err != nil {
		return false
	}

	if _, _, err := CContent(c); err != nil {
		return false
	}
	return true
}

func CTargetAudience(c string) (string, error) {
	if !roughCheck(c) {
		return "", invalidCcodeFormat
	}
	i, _ := byteToInt(pureCcode(c)[0])
	return targets[i], nil
}

func CForm(c string) (string, error) {
	if !roughCheck(c) {
		return "", invalidCcodeFormat
	}
	i, _ := byteToInt(pureCcode(c)[1])
	return forms[i], nil
}

// err, contentTypeName, contentName
func CContent(raw string) (string, string, error) {
	if !roughCheck(raw) {
		return "", "", invalidCcodeFormat
	}
	c := pureCcode(raw)
	s := contents[c[2]]
	category := s[0]
	ci, _ := byteToInt(c[3])
	ci++
	if len(s) <= ci {
		return "", "", fmt.Errorf("Invalid content code (%d)", ci)
	}
	content := s[ci]
	if content == "" {
		return "", "", fmt.Errorf("Invalid content code (%d)", ci)
	}
	return category, content, nil
}
