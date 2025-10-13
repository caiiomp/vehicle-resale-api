package saleApi

import (
	"net/http"

	"github.com/caiiomp/vehicle-resale-api/src/core/responses"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/sale"
	"github.com/gin-gonic/gin"
)

type saleApi struct {
	saleService sale.SaleService
}

func RegisterSaleRoutes(app *gin.Engine, saleService sale.SaleService) {
	service := saleApi{
		saleService: saleService,
	}

	app.GET("/sales", service.search)
}

func (ref *saleApi) search(ctx *gin.Context) {
	sales, err := ref.saleService.Search(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]responses.Sale, len(sales))

	for i, sale := range sales {
		response[i] = responses.SaleFromDomain(sale)
	}

	ctx.JSON(http.StatusOK, response)
}
