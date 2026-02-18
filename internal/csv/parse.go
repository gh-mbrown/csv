package csvparse

import (
	"csv/internal/utils"
	"encoding/csv"
	"log"
	"reflect"
	"strings"
)

func Parse(path string) reflect.Value {
	reader := utils.GetFileReader(path)
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("testing")
	}

	headers := records[0]
	body := records[1:]

	csvT := getCsvType(headers)
	s := reflect.SliceOf(csvT)
	csvV := reflect.MakeSlice(s, 0, 0)

	for _, row := range body {
		v := getCsvValue(csvT, row)
		csvV = reflect.Append(csvV, v)
	}

	return csvV.Elem()
}

func getCsvType(header []string) reflect.Type {
	var fields []reflect.StructField

	for _, col := range header {
		field := reflect.StructField{
			Name: toFieldName(col),
			Type: reflect.TypeFor[string](),
		}

		fields = append(fields, field)
	}

	t := reflect.StructOf(fields)
	return t
}

// Must return Elem() as New() creates a pointer
func getCsvValue(t reflect.Type, record []string) reflect.Value {
	v := reflect.New(t)

	for i, col := range record {
		v.Elem().Field(i).Set(reflect.ValueOf(col))
	}

	return v.Elem()
}

// Formats the field name to be used in a struct. TODO: change this to work with all possibilites
func toFieldName(col string) string {
	parts := strings.Split(col, "_")

	for i, part := range parts {
		if part == "" {
			continue
		}

		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}

	return strings.Join(parts, "")
}
