package token

import "MTDGo/pkg/tokenType"

type Token struct {
	t int
	v string
}

func (tk Token) String() string {
	return tokenType.TypesString[tk.t]
}

func (tk *Token) T() int {
	return tk.t
}

func (tk *Token) SetT(t int) {
	tk.t = t
}

func (tk *Token) V() string {
	return tk.v
}

func (tk *Token) SetV(v string) {
	tk.v = v
}

func NewToken(t int, v string) *Token {
	return &Token{t: t, v: v}
}