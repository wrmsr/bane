package main

import (
	"fmt"
	"strings"
)

func NormalizeWord(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	return s
}

func CheckWord(s string) error {
	if s != NormalizeWord(s) {
		return fmt.Errorf("word not normalized")
	}
	var l rune
	for _, c := range s {
		if !(c >= 'A' && c <= 'Z') && c != ' ' {
			return fmt.Errorf("word contains invalid character: %v", c)
		}
		if l == ' ' && c == ' ' {
			return fmt.Errorf("word contains consecutive spaces")
		}
		l = c
	}
	return nil
}
