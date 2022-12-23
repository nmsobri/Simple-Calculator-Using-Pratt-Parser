package lexer

import (
	"ExprParser/token"
)

type Lexer struct {
	input  string
	cursor int
}

func New(input string) *Lexer {
	return &Lexer{
		input:  input,
		cursor: 0,
	}
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()

	switch l.readChar() {
	case '+':
		t := l.makeToken(token.PLUS, string(l.readChar()), token.PLUSMINUS)
		l.cursor++
		return t

	case '-':
		t := l.makeToken(token.MINUS, string(l.readChar()), token.PLUSMINUS)
		l.cursor++
		return t

	case '/':
		t := l.makeToken(token.SLASH, string(l.readChar()), token.SLASHSTAR)
		l.cursor++
		return t

	case '*':
		t := l.makeToken(token.STAR, string(l.readChar()), token.SLASHSTAR)
		l.cursor++
		return t

	case 0:
		return l.makeToken(token.EOF, "", token.LOWEST)

	default:
		if l.isDigit(l.readChar()) {

			strDigit := ""

			for l.isDigit(l.readChar()) {
				strDigit += string(l.readChar())
				l.cursor++
			}

			return l.makeToken(token.NUMBER, strDigit, token.INTEGER)
		}

		panic("shouldnt arrived here")
	}
}

func (l *Lexer) readChar() byte {
	if l.cursor >= len(l.input) {
		return 0
	}

	return l.input[l.cursor]
}

func (l *Lexer) skipWhiteSpace() {
	for l.isWhiteSpace(l.readChar()) {
		l.cursor++
	}
}

func (*Lexer) isWhiteSpace(ch byte) bool {
	switch ch {
	case ' ', '\t', '\n', '\v', '\f', '\r':
		return true
	default:
		return false
	}
}

func (l *Lexer) makeToken(t token.TokenType, lit string, p int64) token.Token {
	return token.New(t, lit, p)
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
