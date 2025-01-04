package main

import (
	"fmt"
	"strings"
)

func PrintHelp() {
	fmt.Println("alogen : simple markdown html renderer v0.1.0")
	fmt.Println("usage:")
	fmt.Println("    alogen [path] : converts markdown files in path to html")
	fmt.Println("    alogen [file] : converts markdown file to html")
}

func RemoveExt(filename string) string {
	segments := strings.Split(filename, ".")
	if len(segments) == 1 {
		return segments[0]
	}

	ret := ""
	for i := 0; i < len(segments)-1; i++ {
		ret += segments[i]
	}
	return ret
}
