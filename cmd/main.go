package main

import (
	lexer "MTDGo/pkg/lexer"
	"MTDGo/pkg/parser"
	"fmt"
)

func main() {
	input := "# Hello\n### World\n**Hello World**niggers*NiggaWorld*"
	lexer := lexer.NewLexer(input)
	tokens := lexer.Tokenize()
	for _, token := range tokens {
		fmt.Printf("%s %s\n", token, token.V())
	}

	parsed := parser.NewParser(tokens).Parse()

	for _, html := range parsed {
		fmt.Println(html)
	}
}
