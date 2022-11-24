package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ssergomol/Balance-Manager/pkg/apiserver"
)

func main() {
	config, err := apiserver.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	server, err := apiserver.CreateServer(config)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := server.Start(); err != nil {
		logrus.Fatal(err)
	}
}
