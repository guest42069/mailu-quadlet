package main

import (
	"flag"

	"github.com/cyberworm-uk/mailuquadlet"
)

func main() {
	id := flag.String("uuid", "", "optional custom uuid to use for generated")
	compose := flag.String("compose", "docker-compose.yml", "docker-compose.yml file for mailu")
	envfile := flag.String("envfile", "mailu.env", "mailu.env file for mailu")
	flag.Parse()
	m := mailuquadlet.NewMailu(*compose, *envfile)
	if len(*id) > 0 {
		m.Uuid(*id)
	}
	m.Export()
}
