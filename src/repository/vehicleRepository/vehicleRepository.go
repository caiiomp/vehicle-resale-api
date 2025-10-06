package vehicleRepository

import (
	"context"
	"time"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiomp87/vehicle-resale-api/src/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context) ([]entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
}

type vehicleRepository struct {
	collection *mongo.Collection
}

func NewVehicleRepository(collection *mongo.Collection) VehicleRepository {
	return &vehicleRepository{
		collection: collection,
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)
	record.Sold = false

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	created, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id := created.InsertedID.(primitive.ObjectID)

	result := ref.collection.FindOne(ctx, bson.M{"_id": id})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Vehicle
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}

func (ref *vehicleRepository) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var record model.Vehicle
	if err = result.Decode(&record); err != nil {
		return nil, err
	}

	return record.ToDomain(), nil
}

func (ref *vehicleRepository) Search(ctx context.Context) ([]entity.Vehicle, error) {
	cursor, err := ref.collection.Find(ctx, nil)
	if err != nil {
		if err == mongo.ErrNilCursor {
			return nil, nil
		}
		return nil, err
	}

	records := make([]entity.Vehicle, 0)

	for cursor.Next(ctx) {
		var record model.Vehicle
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}

		records = append(records, *record.ToDomain())
	}

	return records, nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)
	record.UpdatedAt = time.Now()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = ref.collection.UpdateOne(ctx, bson.M{"_id": objectID}, record)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Vehicle
	if err = result.Decode(&record); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}
