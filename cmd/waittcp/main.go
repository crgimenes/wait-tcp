package main

import (
	"fmt"

	"crg.eti.br/go/config"
	"crg.eti.br/go/waitTCP"
)

type configData struct {
	Timeout int    `cfg:"timeout" cfgDefault:"1"`
	Host    string `cfg:"host"`
}

func main() {

	cfg := configData{}
	err := config.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	waitTCP.Check(cfg.Host, cfg.Timeout)
}
