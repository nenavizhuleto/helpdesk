package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"application/util"
)


var (
	MONGO       *mongo.Database
	MONGOCLIENT *mongo.Client
)

var BSON_OPTS = options.BSONOptions{
	UseJSONStructTags: true,
}

var COLLECTION_OPTS = options.CollectionOptions{
	BSONOptions: &BSON_OPTS,
}

func MustDisconnectMongo() {
	if err := MONGO.Drop(context.TODO()); err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return MONGO.Collection(name, &COLLECTION_OPTS)
}

func MustConnectMongo(dbname string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(util.MustGetEnvVar("MONGODB")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	log.Println("Connection...")
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	MONGOCLIENT = client
	MONGO = MONGOCLIENT.Database(dbname)
}
