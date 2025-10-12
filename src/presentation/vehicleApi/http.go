package vehicleApi

import (
	"net/http"

	"github.com/caiiomp/vehicle-resale-api/src/core/responses"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/caiiomp/vehicle-resale-api/src/middleware"
	"github.com/gin-gonic/gin"
)

type vehicleApi struct {
	vehicleService vehicle.VehicleService
	authMiddleware middleware.AuthMiddleware
}

func RegisterVehicleRoutes(app *gin.Engine, authMiddleware middleware.AuthMiddleware, vehicleService vehicle.VehicleService) {
	service := vehicleApi{
		vehicleService: vehicleService,
		authMiddleware: authMiddleware,
	}

	app.POST("/vehicles", authMiddleware.Auth, service.create)
	app.GET("/vehicles", service.search)
	app.GET("/vehicles/:vehicle_id", service.get)
	app.PATCH("/vehicles/:vehicle_id", authMiddleware.Auth, service.update)
	app.POST("/vehicles/:vehicle_id/sell", authMiddleware.Auth, service.sell)
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

	response := make([]responses.Vehicle, len(vehicles))

	for i, vehicle := range vehicles {
		response[i] = responses.VehicleFromDomain(vehicle)
	}

	ctx.JSON(http.StatusOK, response)
}

func (ref *vehicleApi) get(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := ref.vehicleService.GetByID(ctx, uri.VehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
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

	vehicle, err := ref.vehicleService.Update(ctx, uri.VehicleID, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}

func (ref *vehicleApi) sell(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := ref.vehicleService.Sell(ctx, uri.VehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}
