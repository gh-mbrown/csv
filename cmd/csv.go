package main

import (
	csvoperations "csv/internal/csv"
	"csv/internal/lexer"
	"flag"
	"fmt"
	"log"
)

func main() {
	var path string
	var query string
	flag.StringVar(&path, "file", "", "file name of the csv file")
	flag.StringVar(&query, "query", "", "query to execute on the file")

	flag.Parse()

	records := csvoperations.LoadCSV(path)

	l := lexer.New(query)
	tokens := l.AllTokens()
	fmt.Println(tokens)

	si := l.GetKeywordIndex(lexer.SELECT)
	// wi := l.GetKeywordIndex(lexer.WHERE)
	// soi := l.GetKeywordIndex(lexer.SORT)
	// li := l.GetKeywordIndex(lexer.LIMIT)

	s_tokens := lexer.NotKeywords(tokens, si+1)

	var columns []string
	for _, v := range s_tokens {
		columns = append(columns, v.Literal)
	}

	stages := []csvoperations.Stage{}

	stages = append(stages, csvoperations.Select{Columns: columns})

	pipeline := csvoperations.Pipeline{Stages: stages}

	records, err := pipeline.Run(records)
	if err != nil {
		log.Fatal("testing")
	}

	// stages := []csvoperations.Stage{}

	// stages = append(stages, csvoperations.Select{Columns: []string{"page", "visits"}})
	// stages = append(stages, csvoperations.Where{Predicate: func(r csvoperations.Record) bool {
	// 	num, err := strconv.Atoi(r["visits"])
	// 	if err != nil {
	// 		log.Fatal("test")
	// 	}
	//
	// 	return num > 1000
	// }})
	// stages = append(stages, csvoperations.Sort{Column: "visits", Desc: false})
	// stages = append(stages, csvoperations.Limit{Count: 2})
	//
	// pipeline := csvoperations.Pipeline{Stages: stages}
	//
	// records, err := pipeline.Run(records)
	// if err != nil {
	// 	log.Fatal("testing")
	// }

	fmt.Println(records)
}
