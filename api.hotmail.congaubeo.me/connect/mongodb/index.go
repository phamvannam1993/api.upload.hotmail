package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// index ...
func index() {
	// Facebook page
	facebookPageIndexes := []mongo.IndexModel{
		newIndex("pageId"),
		newIndex("active"),
		newIndex("createdAt"),
	}
	process(Collection("facebookPages"), facebookPageIndexes)
}

func process(col *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(time.Minute)
	_, err := col.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		fmt.Printf("Index collection %s err: %v", col.Name(), err)
	}
}

func newIndex(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		doc = append(doc, bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		})
	}

	return mongo.IndexModel{Keys: doc}
}
