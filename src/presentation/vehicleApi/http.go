package vehicleApi

import (
	"net/http"

	"github.com/caiomp87/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/gin-gonic/gin"
)

type vehicleApi struct {
	vehicleService vehicle.VehicleService
}

func RegisterVehicleRoutes(app *gin.Engine, vehicleService vehicle.VehicleService) {
	service := vehicleApi{
		vehicleService: vehicleService,
	}

	// POST /vehicle => cadastrar veículo
	// PATCH /vehicle/:vehicle_id => atualizar veículo
	// POST /vehicle/:vehicle_id/buy => comprar veículo
	// GET /vehicle => listar veículos à venda, ordenado por preço, do mais barato ao mais caro
	// GET /vehicle => listar veículos vendidos, ordenado por preço, do mais barato ao mais caro

	app.POST("/vehicle", service.create)
	app.GET("/vehicle", service.search)
	app.GET("/vehicle/:vehicle_id", service.get)
	app.PATCH("/vehicle/:vehicle_id", service.update)
	app.POST("/vehicle/:vehicle_id/buy", service.buy)
}

func (ref *vehicleApi) create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (ref *vehicleApi) search(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (ref *vehicleApi) get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (ref *vehicleApi) update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (ref *vehicleApi) buy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
