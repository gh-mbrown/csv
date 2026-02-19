package csvoperations

type Stage interface {
	Apply(records []Record) ([]Record, error)
}
