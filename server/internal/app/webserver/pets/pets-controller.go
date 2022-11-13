package pets

import (
	"net/http"
	"petclinic-go/server/internal/app/registry"
	"petclinic-go/server/internal/app/services"
	"petclinic-go/server/internal/app/webserver/dtos"

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

		dto := make([]*dtos.PetDTO, len(pets))
		for i, pet := range pets {
			dto[i] = dtos.MapPetToDTO(pet)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"pets": dto,
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

		pet, err := registry.PetsService.CreatePet(ownerId, services.CreateOrUpdatePetParams{
			Name:      request.Name,
			Type:      request.Type,
			BirthDate: request.BirthDate,
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, dtos.MapPetToDTO(pet))
	})

	controller.GET(":petId", func(ctx *gin.Context) {
		ownerId := ctx.Param("ownerId")
		petId := ctx.Param("petId")

		pet, err := registry.PetsService.GetPetByID(ownerId, petId)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dtos.MapPetToDTO(pet))
	})

	controller.GET(":petId/visits", func(ctx *gin.Context) {
		ownerID := ctx.Param("ownerId")
		petID := ctx.Param("petId")

		owner, err := registry.OwnersService.GetByID(ownerID)
		if err != nil {
			ctx.Error(err)
			return
		}

		pet, err := registry.PetsService.GetPetByID(ownerID, petID)
		if err != nil {
			ctx.Error(err)
			return
		}

		visits, err := registry.VisitsService.GetAllForPet(petID)
		if err != nil {
			ctx.Error(err)
			return
		}

		visitDTOs := make([]dtos.VisitDTO, len(visits))
		for i, visit := range visits {
			visitDTOs[i] = *dtos.MapVisitToDTO(visit, owner, pet)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"visits": visitDTOs,
		})
	})

	controller.PUT(":petId", func(ctx *gin.Context) {
		request := &UpdatePetRequest{}
		err := ctx.ShouldBind(request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ownerId := ctx.Param("ownerId")
		petId := ctx.Param("petId")

		pet, err := registry.PetsService.UpdatePet(ownerId, petId, &services.CreateOrUpdatePetParams{
			Name:      request.Name,
			BirthDate: request.BirthDate,
			Type:      request.Type,
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dtos.MapPetToDTO(pet))
	})

	controller.DELETE(":petID", func(ctx *gin.Context) {
		ownerId := ctx.Param("ownerId")
		petId := ctx.Param("petId")

		err := registry.PetsService.DeletePet(ownerId, petId)
		if err != nil {
			ctx.Error(err)
		}

		ctx.Status(http.StatusNoContent)
	})
}
