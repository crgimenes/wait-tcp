package main

import (
	"fmt"
	"os"
	"time"

	"crg.eti.br/go/config"
	"crg.eti.br/go/waitTCP"
)

type configData struct {
	Timeout time.Duration `cfg:"timeout" cfgDefault:"1"`
	Host    string        `cfg:"host" cfgRequired:"true"`
	Verbose bool          `cfg:"verbose" cfgDefault:"false"`
}

var cfg = configData{}

func logf(format string, args ...any) {
	if cfg.Verbose {
		fmt.Printf(format, args...)
	}
}

func main() {
	err := config.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	timeout := cfg.Timeout * time.Second

	logf("trying to connect to %q with timeout %s\n", cfg.Host, timeout)
	if waitTCP.Check(cfg.Host, timeout) {
		logf("timeout\n")
		os.Exit(1)
	}
}
