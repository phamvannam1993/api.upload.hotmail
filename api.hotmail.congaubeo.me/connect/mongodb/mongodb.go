package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log.autofarmer.go/config"
	"log.autofarmer.go/util"
)

var (
	db *mongo.Database
)

// Init ...
func Init() {
	var (
		//envVars = config.GetEnv()
		dbURI  = "mongodb://localhost:27017" /*envVars.Database.URI*/
		dbName = "autofarmer"                /*envVars.Database.Name*/
		//dbAuth  = envVars.Database.Auth
	)
	connectOptions := options.ClientOptions{}
	// Set auth if have
	//if dbAuth.Username != "" && dbAuth.Password != "" {
	connectOptions.Auth = &options.Credential{
		/*			AuthMechanism: "",
					AuthSource:    "",
					Username:      "",
					Password:      "",*/
	}
	//}

	// Connect
	//cl, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(dbURI))
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	cl, err := mongo.Connect(context.TODO(), clientOpts)
	//cl, err := NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	if err != nil {
		log.Fatal("Cannot connect to database", dbURI, err)
	}

	// Set data
	db = cl.Database(dbName)

	// Index db
	index()

	if !config.IsTest() {
		util.ConsolePrintServiceSuccess("MongoDB", dbURI+"/"+dbName)
	}
}
