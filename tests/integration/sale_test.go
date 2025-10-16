//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiiomp/vehicle-resale-api/src/core/responses"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/sale"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/caiiomp/vehicle-resale-api/src/middleware"
	"github.com/caiiomp/vehicle-resale-api/src/presentation"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/saleApi"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/vehicleApi"
	"github.com/caiiomp/vehicle-resale-api/src/repository/memory/saleRepository"
	"github.com/caiiomp/vehicle-resale-api/src/repository/memory/vehicleRepository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListSales(t *testing.T) {
	vehicleRepository := vehicleRepository.NewVehicleRepository()
	saleRepository := saleRepository.NewSaleRepository()

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)
	saleService := sale.NewSaleService(saleRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	vehicleApi.RegisterVehicleRoutes(app, middleware.AuthMiddleware{}, vehicleService)
	saleApi.RegisterSaleRoutes(app, saleService)

	payload := map[string]any{
		"brand": "Ford",
		"model": "Ka",
		"year":  2022,
		"color": "Preto",
		"price": 50000,
	}

	rawPayload, _ := json.Marshal(payload)
	body := bytes.NewReader(rawPayload)

	req, _ := http.NewRequest(http.MethodPost, "/vehicles", body)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var response responses.Vehicle
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	vehicleID := response.ID

	req, _ = http.NewRequest(http.MethodPost, "/vehicles/"+vehicleID+"/buy", nil)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	req, _ = http.NewRequest(http.MethodGet, "/sales", nil)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var salesResponse []responses.Sale
	err = json.Unmarshal(resp.Body.Bytes(), &salesResponse)
	require.NoError(t, err)

	assert.Equal(t, vehicleID, salesResponse[0].VehicleID)
	assert.Equal(t, float64(50000), salesResponse[0].Price)
	assert.NotNil(t, salesResponse[0].SoldAt)
}
