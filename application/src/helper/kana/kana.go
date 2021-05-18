package kana

import (
	"unicode"
)

const (
	// http://www.unicodemap.org/range/62/Hiragana/
	hiraganaLo = 0x3041 // ぁ

	// http://www.unicodemap.org/range/63/Katakana/
	katakanaLo = 0x30a1 // ァ

	codeDiff = katakanaLo - hiraganaLo
)

// HiraganaToKatakana はひらがなをカタカナに変換する。
func HiraganaToKatakana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Hiragana):
			dst[i] = r + codeDiff
		default:
			dst[i] = r
		}
	}
	return string(dst)
}

// KatakanaToHiragana はカタカナをひらがなに変換する。
func KatakanaToHiragana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Katakana):
			dst[i] = r - codeDiff
		default:
			dst[i] = r
		}
	}
	return string(dst)
}