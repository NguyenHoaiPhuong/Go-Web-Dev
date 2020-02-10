package main

import (
	"Go-Web-Dev/087_Mongo-Driver/model"
	"Go-Web-Dev/087_Mongo-Driver/utils"
	"context"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	_ "github.com/mongodb/mongo-go-driver/mongo/options"
)

// TrainerTest : function
func TrainerTest(collection *mongo.Collection) {
	ash := model.Trainer{Name: "Ash", Age: 10, City: "Pallet Town"}
	InsertOneDocument(collection, &ash)

	misty := model.Trainer{Name: "Misty", Age: 10, City: "Cerulean City"}
	brock := model.Trainer{Name: "Brock", Age: 15, City: "Pewter City"}
	trainers := model.Trainers{&misty, &brock}
	toInsert := trainers.ConvertToInterfaceSlice()

	InsertManyDocuments(collection, toInsert)
}

func mongoNow() bson.JavaScript {
    return bson.JavaScript{
      // place your function in here in string
      Code: "(new Date()).ISODate('YYYY-MM-DD hh:mm:ss')"
    }
}

func JavaScripTest(collection *mongo.Collection) {
	err := collection.InsertOne(
		struct{LastSeen interface{}}
		{  
			LastSeen: mongoNow() 
		}
	)
}

func main() {
	client := ConnectToMongoDB()
	dbName := "test"
	colName := "trainers"
	collection := ConnectToCollection(client, dbName, colName)

	// TrainerTest(collection)
	JavaScripTest(collection)

	CloseMongoDBConnection(client)
}

// ConnectToMongoDB : function
func ConnectToMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	utils.CheckError(err)

	err = client.Ping(context.TODO(), nil)
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connected to MongoDB!"))
	return client
}

// CloseMongoDBConnection : function
func CloseMongoDBConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connection to MongoDB closed."))
}

// ConnectToCollection : function
func ConnectToCollection(client *mongo.Client, dbName string, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

// InsertOneDocument : function
func InsertOneDocument(collection *mongo.Collection, document interface{}) {
	insertResult, err := collection.InsertOne(context.TODO(), document)
	utils.CheckError(err)
	fmt.Println(aurora.Green("Inserted a single document: "), aurora.Green(insertResult.InsertedID))
}

// InsertManyDocuments : function
func InsertManyDocuments(collection *mongo.Collection, documents []interface{}) {
	insertManyResult, err := collection.InsertMany(context.TODO(), documents)
	utils.CheckError(err)
	fmt.Println(aurora.Green("Inserted many documents: "), aurora.Green(insertManyResult.InsertedIDs))
}
