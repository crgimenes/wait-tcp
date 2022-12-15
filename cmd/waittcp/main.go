package main

import (
	"fmt"
	"time"

	"crg.eti.br/go/config"
	"crg.eti.br/go/waitTCP"
)

type configData struct {
	Timeout time.Duration `cfg:"timeout" cfgDefault:"1"`
	Host    string        `cfg:"host"`
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
