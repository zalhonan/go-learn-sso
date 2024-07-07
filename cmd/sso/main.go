package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {

	// TODO: config
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: logger

	// TODO: app

	// TODO: run grpc
}
