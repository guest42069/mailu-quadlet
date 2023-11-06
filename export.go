package mailuquadlet

import (
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/compose-spec/compose-go/types"
	"github.com/joho/godotenv"
)

func (m *Mailu) Export() {
	m.volumes()
	m.networks()
	m.containers()
	m.envfile()
}

func (m *Mailu) envfile() {
	envfile := m.uuid + ".env"
	godotenv.Write(m.env, envfile)
	log.Println(envfile)
}

func (m *Mailu) volumes() {
	volumes := map[string]bool{}
	for i := range m.project.Services {
		for j := range m.project.Services[i].Volumes {
			m.project.Services[i].Volumes[j].Source = m.uuid + "-" + path.Base(m.project.Services[i].Volumes[j].Source) + ".volume"
			volumes[m.project.Services[i].Volumes[j].Source] = true
		}
	}
	for volume := range volumes {
		file, err := os.OpenFile(volume, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		vtmpl, err := template.New("volume").Parse(volumeTmpl)
		if err != nil {
			panic(err)
		}
		err = vtmpl.Execute(file, m)
		if err != nil {
			panic(err)
		}
		log.Println(volume)
	}
}

func (m *Mailu) networks() {
	for i := range m.project.Networks {
		network := m.uuid + "-" + i + ".network"
		file, err := os.OpenFile(network, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		ntmpl, err := template.New("network").Parse(networkTmpl)
		if err != nil {
			panic(err)
		}
		err = ntmpl.Execute(file, m.project.Networks[i])
		if err != nil {
			panic(err)
		}
		log.Println(network)
	}
}

func (m *Mailu) containers() {
	for i := range m.project.Services {
		m.project.Services[i].Annotations = types.Mapping{
			"uuid": m.uuid,
		}
		if m.project.Services[i].Name == "resolver" {
			m.project.Services[i].Annotations["ip"] = m.project.Services[i].Networks["default"].Ipv4Address
		}
		if m.project.Services[i].Name == "redis" {
			m.project.Services[i].Image = "docker.io/library" + m.project.Services[i].Image
		}
		container := m.uuid + "-" + m.project.Services[i].Name + ".container"
		file, err := os.OpenFile(container, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		ctmpl, err := template.New("container").Funcs(map[string]any{"contains": strings.Contains}).Parse(containerTmpl)
		if err != nil {
			panic(err)
		}
		err = ctmpl.Execute(file, m.project.Services[i])
		if err != nil {
			panic(err)
		}
		log.Println(container)
	}
}
