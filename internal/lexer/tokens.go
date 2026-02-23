package lexer

import (
	"strings"
)

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	// Speical Tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Ident
	IDENT = "IDENT"

	// Keywords
	SELECT = "SELECT"
	WHERE  = "WHERE"
	SORT   = "SORT"
	LIMIT  = "LIMIT"
	COLUMN = "COLUMN"

	// Misc Characters
	EQUAL         = "EQUAL"
	DOESNOTEQUAL  = "DOESNOTEQUAL"
	LESSTHAN      = "LESSTHAN"
	MORETHAN      = "MORETHAN"
	LESSTHANEQUAL = "LESSTHANEQUAL"
	MORETHANEQUAL = "MORETHANEQUAL"

	STRING = "STRING"
	INT    = "INT"
	FLOAT  = "FLOAT"
)

var keywords = map[string]TokenType{
	"SELECT": SELECT,
	"WHERE":  WHERE,
	"SORT":   SORT,
	"LIMIT":  LIMIT,
}

var operations = map[string]TokenType{
	"==": EQUAL,
	"!=": DOESNOTEQUAL,
	"<":  LESSTHAN,
	">":  MORETHAN,
	"<=": LESSTHANEQUAL,
	">=": MORETHANEQUAL,
}

func LookupKeyword(ident string) TokenType {
	lower := strings.ToUpper(ident)

	if t, ok := keywords[lower]; ok {
		return t
	}

	return IDENT
}

func LookupOperator(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	}

	return ILLEGAL
}

func NotKeywords(tokens []Token, start int) []Token {
	var ts []Token
	s := start

	for s < len(tokens) {
		if _, ok := keywords[string(tokens[s].Type)]; !ok {
			ts = append(ts, tokens[s])
		} else {
			break
		}

		s++
	}

	return ts
}
