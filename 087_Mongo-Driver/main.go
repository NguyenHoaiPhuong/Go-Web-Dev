package main

import (
	"Go-Web-Dev/087_Mongo-Driver/model"
	"Go-Web-Dev/087_Mongo-Driver/utils"
	"context"
	"fmt"
	_ "time"

	"github.com/logrusorgru/aurora"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	_ "github.com/mongodb/mongo-go-driver/mongo/options"
)

func main() {
	client := ConnectToMongoDB()
	dbName := "test"
	colName := "trainers"
	collection := ConnectToCollection(client, dbName, colName)

	ash := model.Trainer{Name: "Ash", Age: 10, City: "Pallet Town"}
	InsertOneDocument(collection, &ash)

	misty := model.Trainer{Name: "Misty", Age: 10, City: "Cerulean City"}
	brock := model.Trainer{Name: "Brock", Age: 15, City: "Pewter City"}
	trainers := model.Trainers{&misty, &brock}
	toInsert := trainers.ConvertToInterfaceSlice()
	InsertManyDocuments(collection, toInsert)

	filter := bson.D{
		{"Name", "Ash"},
	}
	result := collection.FindOne(context.TODO(), filter)
	fmt.Printf("%v\n", result)

	CloseMongoDBConnection(client)
}

func ConnectToMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	utils.CheckError(err)

	err = client.Ping(context.TODO(), nil)
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connected to MongoDB!"))
	return client
}

func CloseMongoDBConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connection to MongoDB closed."))
}

func ConnectToCollection(client *mongo.Client, dbName string, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

func InsertOneDocument(collection *mongo.Collection, document interface{}) {
	insertResult, err := collection.InsertOne(context.TODO(), document)
	utils.CheckError(err)
	fmt.Println(aurora.Green("Inserted a single document: "), aurora.Green(insertResult.InsertedID))
}

func InsertManyDocuments(collection *mongo.Collection, documents []interface{}) {
	insertManyResult, err := collection.InsertMany(context.TODO(), documents)
	utils.CheckError(err)
	fmt.Println(aurora.Green("Inserted many documents: "), aurora.Green(insertManyResult.InsertedIDs))
}

func FindOneDocument(collection *mongo.Collection, filter bson.D) (document interface{}) {
	err := collection.FindOne(context.TODO(), filter).Decode(&document)
	utils.CheckError(err)

	fmt.Printf("Found a single document: %+v\n", document)

	return
}
