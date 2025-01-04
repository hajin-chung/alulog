package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		PrintHelp()
		return
	}

	arg := args[1]
	target_info, err := os.Stat(arg)
	if err != nil {
		fmt.Printf("error on reading target info\n%s\n", err)
		panic(err)
	}

	if target_info.IsDir() {
		convert_dir(arg)
	} else {
		convert_file(arg)
	}
}
