package registry

import (
	"petclinic-go/server/internal/app/repository"
	"petclinic-go/server/internal/app/services"
)

type Registry struct {
	OwnersService *services.OwnersService
	PetsService   *services.PetsService
	VisitsService *services.VisitsService
}

func NewRegistry() *Registry {
	ownersRepository := repository.NewOwnersRepository()
	visitsRepository := repository.NewVisitsRepository()

	registry := &Registry{
		OwnersService: services.NewOwnersService(ownersRepository),
		PetsService:   services.NewPetsService(ownersRepository),
		VisitsService: services.NewVisitsService(ownersRepository, visitsRepository),
	}

	return registry
}
