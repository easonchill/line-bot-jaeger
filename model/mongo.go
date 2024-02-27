package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"line-bot-jaeger/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient interface {
	Collection(name string) *mongo.Collection
}

type mongoClient struct {
	client *mongo.Client
	dbName string
}

func (m *mongoClient) Collection(name string) *mongo.Collection {
	return m.client.Database(m.dbName).Collection(name)
}

func NewMongoClient(client *mongo.Client, dbName string) MongoDBClient {
	return &mongoClient{
		client: client,
		dbName: dbName,
	}
}

func InitMongoDB(cfg *config.MongoDB) MongoDBClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin", cfg.UserName, cfg.Password, cfg.Host, cfg.Database)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	return NewMongoClient(client, cfg.Database)
}
