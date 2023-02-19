package mongoDb

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mongoUri string

var mongoClient *mongo.Client

func InitDb() {
	mongoUri = viper.GetString("database_mongo.uri")
	if mongoUri == "" {
		panic("You must set your 'MONGODB_URI' environmental variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri).SetConnectTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	mongoClient = client
}

func GetMongoDb() *mongo.Database {
	dbName := viper.GetString("database_mongo.db_name")
	db := mongoClient.Database(dbName)
	return db
}
