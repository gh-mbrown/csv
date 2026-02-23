package csvoperations

type Select struct {
	Columns []string
}

func (s Select) Apply(records []Record) ([]Record, error) {
	var result []Record

	for _, record := range records {
		r := make(Record)

		for _, col := range s.Columns {
			if val, ok := record[col]; ok {
				r[col] = val
			}
		}

		result = append(result, r)
	}

	return result, nil
}
