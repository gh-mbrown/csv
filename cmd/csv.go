package main

import (
	csvparse "csv/internal/csv"
	"flag"
)

func main() {
	var path string
	flag.StringVar(&path, "file", "", "file name of the csv file")

	flag.Parse()

	csv := csvparse.Parse(path)
}
