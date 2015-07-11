package cellphone

import (
	. "strconv"
	. "strings"
)

var keys = [8]string{"ABC", "DEF", "GHI", "JKL", "MNO", "PQRS", "TUV", "WXYZ"}

type Cellphone struct{}

func NewCellphone() *Cellphone {
	return &Cellphone{}
}

func (cellphone *Cellphone) convertCharToCode(character rune) string {
	if character == ' ' {
		return "0"
	}

	for idx, key := range keys {
		if ContainsRune(key, character) {
			keystr := Itoa(idx + 2)
			return Repeat(keystr, IndexRune(key, character)+1)
		}
	}
	return ""
}

func (cellphone *Cellphone) isSameGroup(characterA rune, characterB rune) bool {
	for _, key := range keys {
		if ContainsRune(key, characterA) && ContainsRune(key, characterB) {
			return true
		}
	}
	return false
}

func (cellphone *Cellphone) ConvertToKeypadCode(text string) string {
	out := ""
	var beforeCharacter rune
	for _, character := range text {
		if cellphone.isSameGroup(beforeCharacter, character) {
			out += "_"
		}
		out += cellphone.convertCharToCode(character)
		beforeCharacter = character
	}
	return out
}
