package mailuquadlet

import (
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func NewMailu(compose, env string) *Mailu {
	m := &Mailu{}
	m.Init(compose, env)
	return m
}

func (m *Mailu) Init(compose, env string) {
	m.uuid = uuid.NewString()
	envMap, err := godotenv.Read(env)
	if err != nil {
		panic(err)
	}
	m.env = envMap
	project, err := loader.Load(types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{{Filename: compose}},
		Environment: m.env,
	})
	if err != nil {
		panic(err)
	}
	m.project = project
}
