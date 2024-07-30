package mongo

import (
	"context"
	envconfig "learnGin/src/common/envConfig"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var DB *mongo.Database

func LoadMongoDB() {
	if err := ConnectToMongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
	log.Fatalln("Connected to mongodb")
}

// Our implementation logic for connecting to MongoDB
func ConnectToMongodb() error {
	mongoURL := envconfig.GetEnv("MONGO_DB_URI")
	if mongoURL == "" {
		log.Fatal("Mongo URL not found")
		return nil
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	DB = client.Database(envconfig.GetEnv("MONGO_DB_NAME"))
	return err
}
