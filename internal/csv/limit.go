package csvoperations

type Limit struct {
	Count int
}

func (l Limit) Apply(records []Record) ([]Record, error) {
	return records[0:l.Count], nil
}
