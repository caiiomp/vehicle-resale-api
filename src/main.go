package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/sale"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/caiiomp/vehicle-resale-api/src/middleware"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/saleApi"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/vehicleApi"
	"github.com/caiiomp/vehicle-resale-api/src/repository/saleRepository"
	"github.com/caiiomp/vehicle-resale-api/src/repository/vehicleRepository"
)

func main() {
	var (
		mongoURI      = os.Getenv("MONGO_URI")
		mongoDatabase = os.Getenv("MONGO_DATABASE")

		jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("could not initialize mongodb client: %v", err)
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	vehiclesCollection := mongoClient.Database(mongoDatabase).Collection("vehicles")
	salesCollection := mongoClient.Database(mongoDatabase).Collection("sales")

	vehicleRepository := vehicleRepository.NewVehicleRepository(vehiclesCollection)
	saleRepository := saleRepository.NewSaleRepository(salesCollection)

	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)
	saleService := sale.NewSaleService(saleRepository)

	authMiddleware := middleware.NewAuthMiddleware(jwtSecretKey)

	authMiddleware := middleware.NewAuthMiddleware(jwtSecretKey)

	app := gin.Default()

	vehicleApi.RegisterVehicleRoutes(app, authMiddleware, vehicleService)
	saleApi.RegisterSaleRoutes(app, saleService)

	if err = app.Run(":4000"); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
