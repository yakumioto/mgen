package mgen

import "strings"

func SnakeString(str string) string {
	data := make([]byte, 0, len(str)*2)
	noIndex := false
	length := len(str)

	for i := 0; i < length; i++ {
		char := str[i]

		if i > 0 && char >= 'A' && char <= 'Z' && noIndex {
			if i < length-1 {
				lastChar := str[i+1]
				if lastChar >= 'a' && lastChar <= 'z' {
					data = append(data, '_')
				}
			}
		}

		if char != '_' {
			noIndex = true
		}
		data = append(data, char)
	}

	return strings.ToLower(string(data[:]))
}
