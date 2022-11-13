package webserver

import (
	"net/http"
	"petclinic-go/server/internal/app/registry"
	"petclinic-go/server/internal/app/services"

	"github.com/gin-gonic/gin"
)

func PetsController(router *gin.RouterGroup, registry *registry.Registry) {
	controller := router.Group("owners/:ownerId/pets")

	controller.GET("", func(ctx *gin.Context) {
		ownerId := ctx.Param("ownerId")

		pets, err := registry.PetsService.GetPetsForOwnerID(ownerId)
		if err != nil {
			ctx.Error(err)
			return
		}

		dtos := make([]*PetDTO, len(pets))
		for i, pet := range pets {
			dtos[i] = mapPetToDTO(pet)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"pets": dtos,
		})
	})

	controller.POST("", func(ctx *gin.Context) {
		request := &CreatePetRequest{}
		err := ctx.ShouldBind(request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ownerId := ctx.Param("ownerId")

		pet, err := registry.PetsService.CreatePet(ownerId, services.CreatePetParams{
			Name:      request.Name,
			Type:      request.Type,
			BirthDate: request.BirthDate,
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, mapPetToDTO(pet))
	})
}
