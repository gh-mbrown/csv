package csvoperations

import "log"

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
