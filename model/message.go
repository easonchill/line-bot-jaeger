package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionMessage *mongo.Collection

type Message struct {
	UserID    string `json:"userID" bson:"UserID"`
	Text      string `json:"text" bson:"textContent"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}

func NewMessage(userid, text string, timestamp int64) Message {
	return Message{UserID: userid,
		Text:      text,
		Timestamp: timestamp}
}

func InitCollection(client MongoDBClient, name string) {
	collectionMessage = client.Collection(name)
}

func InsertMessage(ctx context.Context, msg Message) (any, error) {
	res, err := collectionMessage.InsertOne(ctx, bson.D{{"UserID", msg.UserID},
		{"textContent", msg.Text},
		{"timestamp", msg.Timestamp},
	})

	return res.InsertedID, err
}

func GetAllMessage(ctx context.Context) []Message {
	cur, err := collectionMessage.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var ms []Message

	for cur.Next(ctx) {
		var result Message
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		ms = append(ms, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return ms
}

func GetUserMessage(ctx context.Context, userid string) []Message {
	cur, err := collectionMessage.Find(ctx, bson.D{{"UserID", userid}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var ms []Message

	for cur.Next(ctx) {
		var result Message
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		ms = append(ms, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return ms
}
