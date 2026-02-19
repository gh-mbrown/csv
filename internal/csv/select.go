package csvoperations

type Select struct {
	Columns []string
}

func (s Select) Apply(records []Record) ([]Record, error) {
	var result []Record

	for _, record := range records {
		r := make(Record)

		for _, col := range s.Columns {
			r[col] = record[col]
		}

		result = append(result, r)
	}

	return result, nil
}
