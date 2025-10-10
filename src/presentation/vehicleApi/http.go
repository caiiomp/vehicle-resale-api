package vehicleApi

import (
	"net/http"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-api/src/core/responses"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/gin-gonic/gin"
)

type vehicleApi struct {
	vehicleService vehicle.VehicleService
}

func RegisterVehicleRoutes(app *gin.Engine, vehicleService vehicle.VehicleService) {
	service := vehicleApi{
		vehicleService: vehicleService,
	}

	app.POST("/vehicle", service.create)
	app.GET("/vehicle", service.search)
	app.GET("/vehicle/:vehicle_id", service.get)
	app.PATCH("/vehicle/:vehicle_id", service.update)
	app.POST("/vehicle/:vehicle_id/buy", service.buy)
}

func (ref *vehicleApi) create(ctx *gin.Context) {
	var request createVehicleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := ref.vehicleService.Create(ctx, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusCreated, response)
}

func (ref *vehicleApi) search(ctx *gin.Context) {
	var query vehicleQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicles, err := ref.vehicleService.Search(ctx, query.IsSold)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]vehicleResponse, len(vehicles))

	for i, vehicle := range vehicles {
		response[i] = vehicleResponseFromDomain(vehicle)
	}

	ctx.JSON(http.StatusOK, response)
}

func (ref *vehicleApi) get(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := ref.vehicleService.GetByID(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := vehicleResponseFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}

func (ref *vehicleApi) update(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request updateVehicleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := ref.vehicleService.Update(ctx, uri.ID, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := vehicleResponseFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}

func (ref *vehicleApi) buy(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sold bool = true
	vehicleToBuy := entity.Vehicle{
		IsSold: &sold,
	}

	vehicle, err := ref.vehicleService.Update(ctx, uri.ID, vehicleToBuy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := vehicleResponseFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}
