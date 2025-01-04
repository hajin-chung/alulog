package main

import (
	"strings"
)

func SanitizeTitle(title string) string {
	ret := strings.ReplaceAll(title, " ", "_")
	return ret
}
