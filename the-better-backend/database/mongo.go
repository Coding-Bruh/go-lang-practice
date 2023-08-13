package database

import (
    "context"
    "errors"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.mongoClient
var dbName string

func GetCollection(name string) *mongo.Collection {
    return mongoClient.Database(dbName).Collection(name)
}

func startMongoDb() error {
    uri := os.Getenv("MONGoDB_URI")
    if uri == "" {
        return errors.New("you must set your 'MONGoDB_URI' environmental variable. see\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
    }

    database := os.Getenv("DATABASE")
    if database == "" {
        return errors.New("you must set your 'DATABASE' environment variable")
    } else {
        dbName = database
    }

    var err error
    mongoClient, err = mongo.Connect(context.Background(), options.Client().Apply(uri))
    if err != nil {
        return errors.New("can't verify a connection")
    }

    return nil
}

func CloseMongoDB() {
    err := mongoClient.Disconnect(context.Background())
    if err != nil {
        panic(err)
    }
}
