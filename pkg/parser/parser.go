package parser

import (
	"MTDGo/pkg/formats"
	"MTDGo/pkg/token"
	"fmt"
)

type Parser struct {
	tokens []*token.Token
	pos, length int
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{tokens: tokens, length: len(tokens)}
}

func (p *Parser) Parse() (result []string) {
	for _, token := range p.tokens {
		format, text := formats.Format[token.T()], token.V()
		if text == "" {
			result = append(result, format)
			continue
		}
		result = append(result, fmt.Sprintf(format, text))
	}
	return
}