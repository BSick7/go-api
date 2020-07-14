package main

import (
	"context"

	"github.com/BSick7/go-api/examples/app1/api"
)

func main() {
	apiServer := app1.Server()
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	apiServer.Launch(8080, cancel)
}
