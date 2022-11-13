package webserver

import (
	"petclinic-go/server/internal/app/registry"
	owners "petclinic-go/server/internal/app/webserver/owners"
	pets "petclinic-go/server/internal/app/webserver/pets"

	"github.com/gin-gonic/gin"
)

func BootstrapWebServer(registry *registry.Registry) {
	r := gin.Default()

	r.Use(ErrorHandler)

	v1 := r.Group("/api/v1")

	owners.OwnersController(v1, registry)
	pets.PetsController(v1, registry)

	r.Run("0.0.0.0:8080")
}
