package mailuquadlet

import (
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"github.com/google/uuid"
)

func NewMailu(compose, env string) *Mailu {
	m := &Mailu{}
	m.Init(compose, env)
	return m
}

func (m *Mailu) Init(compose, env string) {
	m.uuid = uuid.NewString()
	m.env = env
	project, err := loader.Load(types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{{Filename: compose}},
	})
	if err != nil {
		panic(err)
	}
	m.project = project
}
