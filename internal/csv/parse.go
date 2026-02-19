package csvoperations

import (
	"csv/internal/utils"
	"encoding/csv"
	"log"
)

func LoadCSV(path string) []Record {
	reader := utils.GetFileReader(path)
	csvReader := csv.NewReader(reader)

	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("testing")
	}

	var records []Record

	head := lines[0]
	body := lines[1:]

	for _, row := range body {
		record := make(Record)

		for i, cell := range row {
			record[head[i]] = cell
		}

		records = append(records, record)
	}

	return records
}
