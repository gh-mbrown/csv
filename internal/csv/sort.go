package csvoperations

import "sort"

type Sort struct {
	Column string
	Desc   bool
}

func (s Sort) Apply(records []Record) ([]Record, error) {
	sortFunc := func(i int, j int) bool {
		if s.Desc {
			return records[i][s.Column] > records[j][s.Column]
		} else {
			return records[i][s.Column] < records[j][s.Column]
		}
	}
	sort.Slice(records, sortFunc)

	return records, nil
}
