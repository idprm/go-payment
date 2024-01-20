package main

import (
	"log"

	"github.com/idprm/go-payment/src/cmd"
	"github.com/idprm/go-payment/src/config"
)

func main() {

	/**
	 * LOAD CONFIG
	 */
	cfg, err := config.LoadSecret("secret.yaml")
	if err != nil {
		panic(err)
	}

	log.Println(cfg)

	cmd.Execute()
}
