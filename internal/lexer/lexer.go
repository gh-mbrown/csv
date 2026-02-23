package lexer

import (
	"csv/internal/utils"
)

type Lexer struct {
	input  string
	pos    int
	tokens []Token
}

func New(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) peek() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) advance() byte {
	ch := l.input[l.pos]
	l.pos++
	return ch
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && utils.IsWhitespace(l.peek()) {
		l.advance()
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return Token{Type: EOF}
	}

	ch := l.advance()

	if utils.IsDigit(ch) {
		return l.readNumber()
	}

	if utils.IsPeriod(ch) {
		return l.readColumn()
	}

	if utils.IsOperator(ch) {
		return l.readOperator()
	}

	if utils.IsLetter(ch) {
		return l.readIdent()
	}

	return Token{Type: ILLEGAL, Literal: string(ch)}
}

func (l *Lexer) AllTokens() []Token {
	for {
		t := l.NextToken()

		if t.Type == EOF {
			break
		}

		l.tokens = append(l.tokens, t)
	}

	return l.tokens
}

func (l *Lexer) GetKeywordIndex(keyword TokenType) int {
	for i, t := range l.tokens {
		if t.Type == keyword {
			return i
		}
	}

	return -1
}

func (l *Lexer) readNumber() Token {
	start := l.pos - 1
	for l.pos < len(l.input) && utils.IsDigit(l.peek()) {
		l.advance()
	}

	num := l.input[start:l.pos]
	return Token{Type: INT, Literal: num}
}

func (l *Lexer) readIdent() Token {
	start := l.pos - 1
	for l.pos < len(l.input) && utils.IsLetter(l.peek()) {
		l.advance()
	}

	ident := l.input[start:l.pos]

	return Token{Type: LookupKeyword(ident), Literal: ident}
}

func (l *Lexer) readOperator() Token {
	start := l.pos - 1
	for l.pos < len(l.input) && utils.IsOperator(l.peek()) {
		l.advance()
	}

	oper := l.input[start:l.pos]
	return Token{Type: LookupKeyword(oper), Literal: oper}
}

func (l *Lexer) readColumn() Token {
	start := l.pos
	for l.pos < len(l.input) && utils.IsLetter(l.peek()) {
		l.advance()
	}

	column := l.input[start:l.pos]
	return Token{Type: COLUMN, Literal: column}
}
