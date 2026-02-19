package csvoperations

type Where struct {
	Predicate func(Record) bool
}

func (w Where) Apply(records []Record) ([]Record, error) {
	var result []Record

	for _, record := range records {
		if w.Predicate(record) {
			result = append(result, record)
		}
	}

	return result, nil
}
