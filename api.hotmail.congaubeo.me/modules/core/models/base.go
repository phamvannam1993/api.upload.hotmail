package CoreModels

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log.autofarmer.go/config"
	"log.autofarmer.go/connect/mongodb"
)

type MongoDBModels struct {
	Collection string
}
type (
	// AppID ...
	AppID = primitive.ObjectID
	// AppQuery ...
	AppQuery struct {
		Page   int64
		Limit  int64
		Sort   bson.M
		Status string
	}
)

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}
func (model *MongoDBModels) DBCollection() *mongo.Collection {
	if model.Collection == "" {
		log.Println("Empty Collection")
		return nil
	}
	return mongodb.Collection(model.Collection)
}
func (model *MongoDBModels) GetAll(ctx context.Context, query AppQuery) ([]interface{}, int64) {
	var (
		result            = make([]interface{}, 0)
		total       int64 = 0
		wg          sync.WaitGroup
		filter      = bson.M{}
		findOptions = query.GetFindOptionsUsingPage()
	)
	query.AssignStatus(&filter)
	wg.Add(2)
	go func() {
		defer wg.Done()
		result = model.Find(ctx, filter, findOptions)
	}()
	go func() {
		defer wg.Done()
		total = model.Count(ctx, filter)
	}()
	wg.Wait()
	return result, total
}

func (model *MongoDBModels) Count(ctx context.Context, filter interface{}) int64 {
	total, _ := model.DBCollection().CountDocuments(ctx, filter)
	return total
}
func (model *MongoDBModels) Find(ctx context.Context, filter interface{}, opts *options.FindOptions) []interface{} {
	var (
		result = make([]interface{}, 0)
	)

	cursor, err := model.DBCollection().Find(ctx, filter, opts)
	if err != nil {
		return result
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		tempResult := bson.M{}
		err := cursor.Decode(&tempResult)
		if err != nil {
			panic(err)
		}
		result = append(result, tempResult)
	}
	return result
}
func (model *MongoDBModels) FindOne(ctx context.Context, filter interface{}) interface{} {
	var result interface{}
	tempResult := bson.M{}
	err := model.DBCollection().FindOne(ctx, filter).Decode(&tempResult)
	if err == nil {
		obj, _ := json.Marshal(tempResult)
		err = json.Unmarshal(obj, &result)
	}
	return result
}

func (model *MongoDBModels) FindOneByID(ctx context.Context, docID AppID) interface{} {
	var (
		filter = bson.M{"_id": docID}
		result interface{}
	)
	model.DBCollection().FindOne(ctx, filter).Decode(&result)
	return result
}
func (model *MongoDBModels) Insert(ctx context.Context, bsonData interface{}) (interface{}, error) {
	_, err := model.DBCollection().InsertOne(ctx, bsonData)
	return bsonData, err
}
func (model *MongoDBModels) Update(ctx context.Context, filter interface{}, updateData interface{}) error {
	_, err := model.DBCollection().UpdateOne(ctx, filter, updateData)
	return err
}
func (model *MongoDBModels) Delete(ctx context.Context, docID AppID) error {
	var (
		filter = bson.M{"_id": docID}
	)
	_, err := model.DBCollection().DeleteOne(ctx, filter)
	return err
}

func (q AppQuery) GetFindOptionsUsingPage() *options.FindOptions {
	skip := q.Page * q.Limit
	otps := &options.FindOptions{
		Skip:  &skip,
		Limit: &q.Limit,
		Sort:  q.Sort,
	}
	return otps
}
func (q AppQuery) AssignStatus(filter *bson.M) {
	if q.Status == config.QueryStatusActive {
		(*filter)["status.active"] = true
	} else if q.Status == config.QueryStatusInactive {
		(*filter)["status.active"] = false
	}
}
