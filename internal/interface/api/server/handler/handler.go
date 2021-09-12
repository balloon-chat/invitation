// +build !test

package handler

import (
	goEnv "github.com/Netflix/go-env"
	"github.com/balloon/go/invite/env"
	"log"
)

var ClientEntryPoint string

type Environment struct {
	CLIENT_ENTRY_POINT string `env:"CLIENT_ENTRY_POINT"`
}

func init() {
	env.LoadEnv()
	var environment Environment
	_, err := goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables", err)
	}

	if ClientEntryPoint = environment.CLIENT_ENTRY_POINT; ClientEntryPoint == "" {
		log.Fatalln("Environment variable CLIENT_ENTRY_POINT is empty")
	}
}
