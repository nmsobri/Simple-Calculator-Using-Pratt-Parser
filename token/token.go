package token

type TokenType string

const (
	EOF    = "EOF"
	NUMBER = "NUMBER"
	PLUS   = "PLUS"
	MINUS  = "MINUS"
	STAR   = "STAR"
	SLASH  = "SLASH"
)

const (
	LOWEST = iota
	INTEGER
	PLUSMINUS
	SLASHSTAR
)

type Token struct {
	Type       TokenType
	Literal    string
	Precedence int64
}

func New(t TokenType, l string, p int64) Token {
	return Token{
		Type:       t,
		Literal:    l,
		Precedence: p,
	}
}
