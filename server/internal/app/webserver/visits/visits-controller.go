package visits

import (
	"net/http"
	"petclinic-go/server/internal/app/registry"
	"petclinic-go/server/internal/app/services"
	"petclinic-go/server/internal/app/webserver/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func VisitsController(router *gin.RouterGroup, registry *registry.Registry) {
	controller := router.Group("visits")

	controller.GET("", func(ctx *gin.Context) {
		visits, err := registry.VisitsService.GetAll()
		if err != nil {
			ctx.Error(err)
			return
		}

		visitDTOs := make([]dtos.VisitDTO, len(visits))
		for i, visit := range visits {
			owner, err := registry.OwnersService.GetByID(visit.OwnerID)
			if err != nil {
				ctx.Error(err)
				return
			}

			pet, err := registry.PetsService.GetPetByID(visit.OwnerID, visit.PetID)
			if err != nil {
				ctx.Error(err)
				return
			}

			visitDTOs[i] = *dtos.MapVisitToDTO(visit, owner, pet)
			i++
		}

		ctx.JSON(http.StatusOK, gin.H{
			"visits": visitDTOs,
		})
	})

	controller.GET(":visitId", func(ctx *gin.Context) {
		visitId := ctx.Param("visitId")

		visit, err := registry.VisitsService.GetByID(visitId)
		if err != nil {
			ctx.Error(err)
			return
		}

		owner, err := registry.OwnersService.GetByID(visit.OwnerID)
		if err != nil {
			ctx.Error(err)
			return
		}

		pet, err := registry.PetsService.GetPetByID(visit.OwnerID, visit.PetID)
		if err != nil {
			ctx.Error(err)
			return
		}

		dto := dtos.MapVisitToDTO(visit, owner, pet)

		ctx.JSON(http.StatusOK, dto)
	})

	controller.POST("", func(ctx *gin.Context) {
		request := &CreateVisitRequest{}
		err := ctx.Bind(request)

		if err != nil {
			ctx.Error(err)
			return
		}

		date, err := time.ParseInLocation(time.RFC3339, request.Date, time.UTC)
		if err != nil {
			ctx.Error(err)
			return
		}

		owner, err := registry.OwnersService.GetByID(request.OwnerID)
		if err != nil {
			ctx.Error(err)
			return
		}

		pet, err := registry.PetsService.GetPetByID(request.OwnerID, request.PetID)
		if err != nil {
			ctx.Error(err)
			return
		}

		visit, err := registry.VisitsService.Create(&services.CreateOrUpdateVisitParams{
			Date:        date,
			Description: request.Description,

			OwnerID: owner.ID,
			PetID:   pet.ID,
		})
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, dtos.MapVisitToDTO(visit, owner, pet))
	})

	controller.PUT(":visitId", func(ctx *gin.Context) {
		request := &UpdateVisitRequest{}
		err := ctx.Bind(request)

		if err != nil {
			ctx.Error(err)
			return
		}

		visitID := ctx.Param("visitId")

		date, err := time.ParseInLocation(time.RFC3339, request.Date, time.UTC)
		if err != nil {
			ctx.Error(err)
			return
		}

		owner, err := registry.OwnersService.GetByID(request.OwnerID)
		if err != nil {
			ctx.Error(err)
			return
		}

		pet, err := registry.PetsService.GetPetByID(request.OwnerID, request.PetID)
		if err != nil {
			ctx.Error(err)
			return
		}

		visit, err := registry.VisitsService.Update(visitID, &services.CreateOrUpdateVisitParams{
			Date:        date,
			Description: request.Description,

			OwnerID: owner.ID,
			PetID:   pet.ID,
		})
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dtos.MapVisitToDTO(visit, owner, pet))
	})

	controller.DELETE(":visitId", func(ctx *gin.Context) {
		visitID := ctx.Param("visitId")

		err := registry.VisitsService.Delete(visitID)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.Status(http.StatusNoContent)
	})
}
