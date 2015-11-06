package lexer

import (
	"container/list"
	"fmt"
	"regexp"
	"strings"
)

// TokenType enum
type TokenType int

// Token type
const (
	Symbol TokenType = iota
	Number
	Bool
	List
	Proc
	Lambda
	LeftParenthes
	RightParenthes
)

func (t TokenType) String() string {
	switch t {
	case Symbol:
		return "Symbol"
	case Number:
		return "Number"
	case Bool:
		return "Bool"
	case List:
		return "List"
	case Proc:
		return "<Proc>"
	case Lambda:
		return "<Lambda>"
	case LeftParenthes:
		return "("
	case RightParenthes:
		return ")"
	}
	return "<unknown>"
}

// Token struct
type Token struct {
	Type TokenType
	Raw  string
	Val  interface{}
}

func (t *Token) String() string {
	return fmt.Sprintf("%v", t.Val)
}

// Tokenizer func
func Tokenizer(expr string) (tokens *list.List, err error) {
	expr = strings.TrimSpace(expr)
	if len(expr) == 0 {
		return
	}
	tokens = list.New()
	i := 0
	for i < len(expr) {
		if expr[i] == '(' {
			tokens.PushBack(&Token{Type: LeftParenthes, Raw: "(", Val: "("})
			break
		}
		if expr[i] == ')' {
			token := atom(expr[:i])
			tokens.PushBack(&token)
			tokens.PushBack(&Token{Type: RightParenthes, Raw: ")", Val: ")"})
			break
		}
		if expr[i] == ' ' || expr[i] == '\t' {
			token := atom(expr[:i])
			tokens.PushBack(&token)
			break
		}
		i++
	}
	if i == len(expr) {
		token := atom(expr[:i])
		tokens.PushBack(&token)
		return
	}
	tails, err := Tokenizer(expr[i+1:])
	if err != nil {
		return tokens, err
	}
	if tails != nil {
		tokens.PushBackList(tails)
	}
	return
}

var digitRegexp = regexp.MustCompile(`-?\d+`)

func isDigit(expr string) bool {
	return digitRegexp.MatchString(expr)
}

func atom(raw string) Token {
	if raw == "#f" {
		return Token{Type: Bool, Raw: raw, Val: false}
	} else if raw == "#t" {
		return Token{Type: Bool, Raw: raw, Val: true}
	}
	if isDigit(raw) {
		return Token{Type: Number, Raw: raw, Val: raw}
	}
	return Token{Type: Symbol, Raw: raw, Val: raw}
}
