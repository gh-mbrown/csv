package csvoperations

import (
	"csv/internal/lexer"
	"strconv"
)

var Operations = map[lexer.TokenType]func(args ...string) Stage{
	lexer.SELECT: func(args ...string) Stage {
		return Select{Columns: args}
	},
	lexer.LIMIT: func(args ...string) Stage {
		n, _ := strconv.Atoi(args[0])
		return Limit{Count: n}
	},
	lexer.SORT: func(args ...string) Stage {
		desc := len(args) > 1 && args[1] == "DESC"
		return Sort{Column: args[0], Desc: desc}
	},
}
