package lexer

import (
	"MTDGo/pkg/token"
	"MTDGo/pkg/tokenType"
	"log"
	"strings"
	"unicode"
)

const SYMS = "*`#[]()-"

var Headers [6]int = [6]int{
	0: tokenType.H1,
	1: tokenType.H2,
	2: tokenType.H3,
	3: tokenType.H4,
	4: tokenType.H5,
	5: tokenType.H6,
}

var BoldAndItalic [3]int = [3]int{
	0: tokenType.Italic,
	1: tokenType.Bold,
	2: tokenType.BoldItalic,
}

type Lexer struct {
	input string
	tokens []*token.Token
	pos, length int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, length: len(input)}
}

func (l *Lexer) Tokenize() []*token.Token {
	for l.pos <= l.length {
		cur := l.peek(0)
		switch {
		case unicode.IsLetter(cur):
			l.tokenizeText()
		case cur == '#':
			l.tokenizeHeader()
		case cur == '*':
			l.tokenizeBoldAndItalic()
		case cur == '\n':
			l.addToken(tokenType.NewLine, "")
			fallthrough
		default:
			l.next()
		}
	}
	return l.tokens
}

func (l *Lexer) tokenizeHeader() {
	var headerType int
	var buf string
	for {
		cur := l.peek(0)
		if unicode.IsSpace(cur) {
			break
		}
		if cur != '#' {
			log.Fatal("Syntax Error!")
		}

		headerType++
		cur = l.next()
	}
	l.next()
	for {
		cur := l.peek(0)
		if cur == '\n' {
			//l.addToken(tokenType.NewLine, "")
			break
		}
		buf += string(cur)
		cur = l.next()
	}
	l.addToken(Headers[headerType-1], buf)
}

func (l *Lexer) tokenizeBoldAndItalic() {
	var starsCounter int
	var buf string
	for {
		cur := l.peek(0)
		if unicode.IsLetter(cur) {
			break
		}
		if cur != '*' {
			log.Fatal("Syntax Error!")
		}

		starsCounter++
		cur = l.next()
	}
	for {
		cur := l.peek(0)
		if cur == '*' {
			for i := 0; i < starsCounter; i++ {
				l.next()
			}
			break
		}
		buf += string(cur)
		cur = l.next()
	}
	l.addToken(BoldAndItalic[starsCounter-1], buf)
}

func (l *Lexer) tokenizeText() {
	var buf string
	for {
		cur := l.peek(0)
		if strings.ContainsRune(SYMS, cur) {
			break
		}
		if cur == '\n' {
			//l.addToken(tokenType.NewLine, "")
			break
		}
		buf += string(cur)
		cur = l.next()
	}
	l.addToken(tokenType.Text, buf)
}

func (l *Lexer) next() rune {
	l.pos++
	return l.peek(0)
}

func (l *Lexer) peek(relativePosition int) rune {
	pos := l.pos + relativePosition
	if pos >= l.length {
		return '\n'
	}
	return rune(l.input[pos])
}

func (l *Lexer) addToken(t int, v string) {
	l.tokens = append(l.tokens, token.NewToken(t, v))
}






