package main

import (
	"petclinic-go/server/internal/app/registry"
	"petclinic-go/server/internal/app/webserver"
)

func main() {
	registry := registry.NewRegistry()

	webserver.BootstrapWebServer(registry)
}
