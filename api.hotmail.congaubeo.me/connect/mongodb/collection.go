package mongodb

import "go.mongodb.org/mongo-driver/mongo"

func Collection(CollectionName string) *mongo.Collection {
	return db.Collection(CollectionName)
}
