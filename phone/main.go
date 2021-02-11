package main

import (
	"strings"
)

func normalize(phone string) string {
	var sb strings.Builder

	for _, char := range phone {
		if char >= '0' && char <= '9' {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}
