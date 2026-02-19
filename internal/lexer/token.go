package lexer

import (
	"fmt"
	"strings"
)

type Token int

const (
	SELECT Token = iota
	WHERE
	SORT
	LIMIT
	COLUMN
	EQUAL
	DOESNOTEQUAL
	LESSTHAN
	MORETHAN
	LESSTHANEQUAL
	MORETHANEQUAL
	STRING
	INT
	FLOAT
)

func (t Token) String() string {
	switch t {
	case SELECT:
		return "select"
	case WHERE:
		return "where"
	case COLUMN:
		return "column"
	case DOESNOTEQUAL:
		return "doesnotequal"
	case EQUAL:
		return "equal"
	case FLOAT:
		return "float"
	case INT:
		return "int"
	case LESSTHAN:
		return "lessthan"
	case LESSTHANEQUAL:
		return "lessthanequal"
	case LIMIT:
		return "limit"
	case MORETHAN:
		return "morethan"
	case MORETHANEQUAL:
		return "morethanequal"
	case SORT:
		return "sort"
	case STRING:
		return "string"
	default:
		panic(fmt.Sprintf("unexpected lexer.Token: %#v", t))
	}
}

func Tokenize(input string) ([]Token, error) {
	lowered := strings.ToLower(input)

	var aux func(toToken string, index int, tokens []Token) []Token

	aux = func(toToken string, index int, tokens []Token) []Token {
		if len(toToken) <= 0 || toToken == "" {
			return tokens
		}

		if strings.HasPrefix(toToken[index:], "select") {
			t := append(tokens, SELECT)
			i := index + 6
			return aux(toToken[i:], i, t)
		} else if strings.HasPrefix(toToken[index:], "where") {
			t := append(tokens, WHERE)
			i := index + 5
			return aux(toToken[i:], i, t)
		}
	}

	tokens := make([]Token, 0)

	return aux(lowered, 0, tokens), nil

}
