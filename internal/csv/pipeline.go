package csvoperations

import "log"

type Stage interface {
	Apply(records []Record) ([]Record, error)
}

type Pipeline struct {
	Stages []Stage
}

func (p *Pipeline) Run(records []Record) ([]Record, error) {
	var err error
	for _, stage := range p.Stages {
		records, err = stage.Apply(records)
		if err != nil {
			log.Fatal("testing")
		}
	}

	return records, nil
}
