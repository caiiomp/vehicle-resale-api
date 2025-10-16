//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiiomp/vehicle-resale-api/src/core/responses"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/caiiomp/vehicle-resale-api/src/middleware"
	"github.com/caiiomp/vehicle-resale-api/src/presentation"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/vehicleApi"
	"github.com/caiiomp/vehicle-resale-api/src/repository/memory/saleRepository"
	"github.com/caiiomp/vehicle-resale-api/src/repository/memory/vehicleRepository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateVehicle(t *testing.T) {
	vehicleRepository := vehicleRepository.NewVehicleRepository()
	saleRepository := saleRepository.NewSaleRepository()

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	vehicleApi.RegisterVehicleRoutes(app, middleware.AuthMiddleware{}, vehicleService)

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

	assert.Equal(t, "Ford", response.Brand)
	assert.Equal(t, "Ka", response.Model)
	assert.Equal(t, 2022, response.Year)
	assert.Equal(t, "Preto", response.Color)
	assert.Equal(t, float64(50000), response.Price)
	assert.NotNil(t, response.CreatedAt)
	assert.NotNil(t, response.UpdatedAt)
	assert.Nil(t, response.SoldAt)
}

func TestSearchVehicles(t *testing.T) {
	vehicleRepository := vehicleRepository.NewVehicleRepository()
	saleRepository := saleRepository.NewSaleRepository()

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	vehicleApi.RegisterVehicleRoutes(app, middleware.AuthMiddleware{}, vehicleService)

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

	req, _ = http.NewRequest(http.MethodGet, "/vehicles", nil)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var vehicles []responses.Vehicle
	err = json.Unmarshal(resp.Body.Bytes(), &vehicles)
	require.NoError(t, err)

	assert.Equal(t, vehicleID, vehicles[0].ID)
	assert.Equal(t, "Ford", vehicles[0].Brand)
	assert.Equal(t, "Ka", vehicles[0].Model)
	assert.Equal(t, 2022, vehicles[0].Year)
	assert.Equal(t, "Preto", vehicles[0].Color)
	assert.Equal(t, float64(50000), vehicles[0].Price)
	assert.NotNil(t, vehicles[0].CreatedAt)
	assert.NotNil(t, vehicles[0].UpdatedAt)
	assert.Nil(t, vehicles[0].SoldAt)
}

func TestGetVehicle(t *testing.T) {
	vehicleRepository := vehicleRepository.NewVehicleRepository()
	saleRepository := saleRepository.NewSaleRepository()

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	vehicleApi.RegisterVehicleRoutes(app, middleware.AuthMiddleware{}, vehicleService)

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

	req, _ = http.NewRequest(http.MethodGet, "/vehicles/"+vehicleID, nil)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, vehicleID, response.ID)
	assert.Equal(t, "Ford", response.Brand)
	assert.Equal(t, "Ka", response.Model)
	assert.Equal(t, 2022, response.Year)
	assert.Equal(t, "Preto", response.Color)
	assert.Equal(t, float64(50000), response.Price)
	assert.NotNil(t, response.CreatedAt)
	assert.NotNil(t, response.UpdatedAt)
	assert.Nil(t, response.SoldAt)
}

func TestUpdateVehicle(t *testing.T) {
	vehicleRepository := vehicleRepository.NewVehicleRepository()
	saleRepository := saleRepository.NewSaleRepository()

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	vehicleApi.RegisterVehicleRoutes(app, middleware.AuthMiddleware{}, vehicleService)

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
	oldUpdatedAt := response.UpdatedAt

	payload = map[string]any{
		"brand": "Chevrolet",
		"model": "Onix",
		"year":  2023,
		"color": "Branco",
		"price": 60000,
	}

	rawPayload, _ = json.Marshal(payload)
	body = bytes.NewReader(rawPayload)

	req, _ = http.NewRequest(http.MethodPatch, "/vehicles/"+vehicleID, body)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Chevrolet", response.Brand)
	assert.Equal(t, "Onix", response.Model)
	assert.Equal(t, 2023, response.Year)
	assert.Equal(t, "Branco", response.Color)
	assert.Equal(t, float64(60000), response.Price)
	assert.Greater(t, response.UpdatedAt, oldUpdatedAt)
	assert.NotNil(t, response.CreatedAt)
	assert.NotNil(t, response.UpdatedAt)
	assert.Nil(t, response.SoldAt)
}

func TestBuyVehicle(t *testing.T) {
	vehicleRepository := vehicleRepository.NewVehicleRepository()
	saleRepository := saleRepository.NewSaleRepository()

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	vehicleApi.RegisterVehicleRoutes(app, middleware.AuthMiddleware{}, vehicleService)

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

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Ford", response.Brand)
	assert.Equal(t, "Ka", response.Model)
	assert.Equal(t, 2022, response.Year)
	assert.Equal(t, "Preto", response.Color)
	assert.Equal(t, float64(50000), response.Price)
	assert.NotNil(t, response.CreatedAt)
	assert.NotNil(t, response.UpdatedAt)
	assert.NotNil(t, response.SoldAt)
}
