package main

import (
	"fmt"
	"os"
	"yyds/internal/parser"
)

func main() {
	if len(os.Args) != 2 {
		printUsage()
		return
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	m := parser.Parse(string(content))
	m.Play()
}

func printUsage() {
	fmt.Println("Usage of yyds: ")
	fmt.Println("    yyds file_to_play")
}
