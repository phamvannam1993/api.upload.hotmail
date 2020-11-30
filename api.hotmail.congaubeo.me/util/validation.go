package util

import "go.mongodb.org/mongo-driver/bson/primitive"

// ValidationIsObjectID ...
func ValidationIsObjectID(id string) (result primitive.ObjectID, err error) {
	result, err = primitive.ObjectIDFromHex(id)
	return
}
