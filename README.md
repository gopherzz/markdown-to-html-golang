# Markdown To Html Converter

## Simple Example
```go
  package main

import (
	"github.com/gopherzz/MTDGo/pkg/lexer"
	"github.com/gopherzz/MTDGo/pkg/parser"
	"fmt"
)

func main() {
	input := "# Hello\n### World\n*Hello World*I love**Golang**"
	tokens := lexer.NewLexer(input).Tokenize()
	parsed := parser.NewParser(tokens).Parse()

	for _, html := range parsed {
		fmt.Println(html)
	}
}

```
