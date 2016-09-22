package str

import (
	"unicode"
	"strings"
)

func UpcaseInitial(str string) string {

	cstr := strings.TrimSpace(str)
	runes := []rune(cstr)

	for i, v := range runes {

		
		return string(unicode.ToUpper(v)) + string(runes[i+1:])

	}

	return ""

}
