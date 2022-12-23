package parser

import (
	"ExprParser/lexer"
	"ExprParser/token"
	"fmt"
	"strconv"
)

type prefixFn func() int64
type infixFn func(val int64) int64

type Parser struct {
	lexer           *lexer.Lexer
	curToken        token.Token
	peekToken       token.Token
	prefixFunctions map[token.TokenType]prefixFn
	infixFunctions  map[token.TokenType]infixFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	p.prefixFunctions = make(map[token.TokenType]prefixFn)
	p.infixFunctions = make(map[token.TokenType]infixFn)

	p.prefixFunctions[token.NUMBER] = p.parseNumber

	p.infixFunctions[token.PLUS] = p.parseInfix
	p.infixFunctions[token.MINUS] = p.parseInfix
	p.infixFunctions[token.STAR] = p.parseInfix
	p.infixFunctions[token.SLASH] = p.parseInfix

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() int64 {
	return p.parseExpression(token.LOWEST)
}

func (p *Parser) parseExpression(precedence int64) int64 {
	prefixFn := p.prefixFunctions[p.curToken.Type]

	if prefixFn == nil {
		fmt.Printf("Missing prefix function for token: %s\n", p.curToken.Type)
		return 0
	}

	left := prefixFn()

	for p.peekPrecedence() > precedence {
		infixFn := p.infixFunctions[p.peekToken.Type]

		if infixFn == nil {
			fmt.Printf("Missing infix function for token: %s\n", p.peekToken.Type)
			return 0
		}

		p.nextToken() // eat infix token
		left = infixFn(left)
	}

	return left
}

func (p *Parser) parseNumber() int64 {
	num, err := strconv.ParseInt(p.curToken.Literal, 10, 64)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return num
}

func (p *Parser) parseInfix(left int64) int64 {
	currentToken := p.curToken
	p.nextToken() // advance to next token

	switch currentToken.Type {
	case token.PLUS:
		return left + p.parseExpression(token.PLUSMINUS)

	case token.MINUS:
		return left - p.parseExpression(token.PLUSMINUS)

	case token.STAR:
		return left * p.parseExpression(token.SLASHSTAR)

	case token.SLASH:
		return left / p.parseExpression(token.SLASHSTAR)

	default:
		panic("shouldnt arrived here")
	}

}

func (p *Parser) peekPrecedence() int64 {
	return p.peekToken.Precedence
}
