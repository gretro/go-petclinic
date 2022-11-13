package registry

import (
	"petclinic-go/server/internal/app/repository"
	"petclinic-go/server/internal/app/services"
)

type Registry struct {
	OwnersService *services.OwnersService
	PetsService   *services.PetsService
}

func NewRegistry() *Registry {
	ownersRepository := repository.NewOwnersRepository()

	registry := &Registry{
		OwnersService: services.NewOwnersService(ownersRepository),
		PetsService:   services.NewPetsService(ownersRepository),
	}

	return registry
}
