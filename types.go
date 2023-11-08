package mailuquadlet

import "github.com/compose-spec/compose-go/types"

type Mailu struct {
	uuid    string
	env     string
	project *types.Project
}

func (m *Mailu) Uuid(id string) {
	m.uuid = id
}
