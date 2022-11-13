package owners

import (
	"net/http"
	"petclinic-go/server/internal/app/registry"
	"petclinic-go/server/internal/app/services"
	"petclinic-go/server/internal/app/webserver/dtos"

	"github.com/gin-gonic/gin"
)

func OwnersController(router *gin.RouterGroup, registry *registry.Registry) {
	controller := router.Group("owners")

	controller.GET("", func(ctx *gin.Context) {
		owners := registry.OwnersService.GetAll()

		ownerDTOs := make([]*dtos.OwnerSummaryDTO, len(owners))
		for i, owner := range owners {
			ownerDTOs[i] = dtos.MapOwnerToSummaryDTO(owner)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"owners": ownerDTOs,
		})
	})

	controller.GET(":ownerId", func(ctx *gin.Context) {
		ownerId, _ := ctx.Params.Get("ownerId")

		owner, error := registry.OwnersService.GetByID(ownerId)

		if error != nil {
			ctx.Error(error)
			return
		}

		dto := dtos.MapOwnerToDTO(owner)

		ctx.JSON(http.StatusOK, dto)
	})

	controller.POST("", func(ctx *gin.Context) {
		request := &CreateAndUpdateOwnerRequest{}
		err := ctx.ShouldBind(request)

		if err != nil {
			ctx.Error(err)
			return
		}

		owner, err := registry.OwnersService.CreateOwner(services.CreateOrUpdateOwnerParams{
			FirstName: request.FirstName,
			LastName:  request.LastName,

			Address:     request.Address,
			City:        request.City,
			PhoneNumber: request.PhoneNumber,
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		dto := dtos.MapOwnerToDTO(owner)

		ctx.JSON(http.StatusCreated, dto)
	})

	controller.PUT(":ownerId", func(ctx *gin.Context) {
		request := &CreateAndUpdateOwnerRequest{}
		err := ctx.ShouldBind(request)

		if err != nil {
			ctx.Error(err)
			return
		}

		ownerID := ctx.Param("ownerId")

		owner, err := registry.OwnersService.UpdateOwner(ownerID, services.CreateOrUpdateOwnerParams{
			FirstName: request.FirstName,
			LastName:  request.LastName,

			Address:     request.Address,
			City:        request.City,
			PhoneNumber: request.PhoneNumber,
		})

		if err != nil {
			ctx.Error(err)
			return
		}

		dto := dtos.MapOwnerToDTO(owner)

		ctx.JSON(http.StatusOK, dto)
	})

	controller.DELETE(":ownerId", func(ctx *gin.Context) {
		ownerId, _ := ctx.Params.Get("ownerId")

		registry.OwnersService.DeleteOwnerByID(ownerId)

		ctx.Status(http.StatusNoContent)
	})
}
