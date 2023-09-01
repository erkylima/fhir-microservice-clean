package dbmongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const OperationTimeOut = 5

func NewCollection(uriConnection, dbName, collectionName string) (*mongo.Collection, error) {
	Ctx, _ := context.WithTimeout(context.TODO(), OperationTimeOut*time.Second)
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uriConnection))
	if err != nil {
		return nil, fmt.Errorf("connection Failed %w", err)
	}
	if err := client.Ping(Ctx, nil); err != nil {
		return nil, fmt.Errorf("connection Failed %w", err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return collection, nil
}
