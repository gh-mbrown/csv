package main

import (
	csvoperations "csv/internal/csv"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	var path string
	flag.StringVar(&path, "file", "", "file name of the csv file")

	flag.Parse()

	records := csvoperations.LoadCSV(path)

	stages := []csvoperations.Stage{}

	stages = append(stages, csvoperations.Select{Columns: []string{"page", "visits"}})
	stages = append(stages, csvoperations.Where{Predicate: func(r csvoperations.Record) bool {
		num, err := strconv.Atoi(r["visits"])
		if err != nil {
			log.Fatal("test")
		}

		return num > 1000
	}})
	stages = append(stages, csvoperations.Sort{Column: "visits", Desc: false})
	stages = append(stages, csvoperations.Limit{Count: 2})

	pipeline := csvoperations.Pipeline{Stages: stages}

	records, err := pipeline.Run(records)
	if err != nil {
		log.Fatal("testing")
	}

	fmt.Println(records)
}
