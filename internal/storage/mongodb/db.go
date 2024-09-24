package mongodb

import (
	"context"
	"go-fiber-hex-arch/internal/adapter/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

type DB struct {
	db         *mongo.Database
	collection string
}

func New(db *mongo.Database, collection string) (*DB, error) {
	return &DB{
		db:         db,
		collection: collection,
	}, nil
}

func NewDB(cfg *config.MONGO) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	slog.Info("Connected to mongodb")

	return client, nil
}
