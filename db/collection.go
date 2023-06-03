package db

import "go.mongodb.org/mongo-driver/mongo"

func Default(client *mongo.Client, collectionName string)*mongo.Collection{
	return client.Database("mail-app").Collection(collectionName)
}