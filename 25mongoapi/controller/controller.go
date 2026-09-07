package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aviraltripathi25_db_user:EZTIld2aP6s4gXhz@cluster0.pzzyhjy.mongodb.net/?appName=Cluster0"
const dbName = "netmirror"
const colName = "watchlist"

// imp stuff
var collection *mongo.Collection

//connect with mongodb

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}
