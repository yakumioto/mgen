package mgen

import "strings"

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	if num == 2 {
		if s == "ID" {
			return "id"
		}
	}
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			if s[i] == 'I' && s[i+1] == 'D' {
				data = append(data, '_', s[i], s[i+1])
				i++
				continue
			}
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
