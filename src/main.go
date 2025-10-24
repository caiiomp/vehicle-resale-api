package main

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/sale"
	"github.com/caiiomp/vehicle-resale-api/src/core/useCases/vehicle"
	"github.com/caiiomp/vehicle-resale-api/src/presentation"

	_ "github.com/caiiomp/vehicle-resale-api/src/docs"
	"github.com/caiiomp/vehicle-resale-api/src/middleware"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/saleApi"
	"github.com/caiiomp/vehicle-resale-api/src/presentation/vehicleApi"
	"github.com/caiiomp/vehicle-resale-api/src/repository/mongodb/saleRepository"
	"github.com/caiiomp/vehicle-resale-api/src/repository/mongodb/vehicleRepository"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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

	app := presentation.SetupServer()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vehicleApi.RegisterVehicleRoutes(app, authMiddleware, vehicleService)
	saleApi.RegisterSaleRoutes(app, saleService)

	if err = app.Run(":8080"); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
