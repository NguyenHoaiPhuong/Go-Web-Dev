package main

import (
	"Go-Web-Dev/087_Mongo-Driver/model"
	"context"
	"strconv"
	"testing"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func TestBulkWrite(t *testing.T) {
	client := ConnectToMongoDB()
	dbName := "test"
	colName := "trainers"
	collection := ConnectToCollection(client, dbName, colName)

	ash := model.Trainer{Name: "Ash", Age: 10, City: "Pallet Town"}
	misty := model.Trainer{Name: "Misty", Age: 10, City: "Cerulean City"}
	brock := model.Trainer{Name: "Brock", Age: 15, City: "Pewter City"}
	trainers := model.Trainers{&ash, &misty, &brock}

	models := make([]mongo.WriteModel, len(trainers))
	for i, trainer := range trainers {
		var model mongo.WriteModel
		iom := mongo.NewInsertOneModel()
		if trainer != nil {
			iom = iom.SetDocument(trainer)
		}
		model = iom
		models[i] = model
	}
	res, err := collection.BulkWrite(context.TODO(), models)
	if err != nil {
		t.Error(err)
	}
	if res.InsertedCount != int64(len(trainers)) {
		t.Error("Number of items written onto MongoDB is wrong")
	}

	CloseMongoDBConnection(client)
}

func CreateCOGS() []model.COGS {
	cogsSlice := make([]model.COGS, 68*3)
	for i := 0; i < 68; i++ {
		for j := 0; j < 3; j++ {
			cogsSlice[i*3+j].ProductID = "T" + strconv.Itoa(i+1) + "_Product" + strconv.Itoa(j+1)
			cogsSlice[i*3+j].COGSCLPPerKg = 1.0
			cogsSlice[i*3+j].ScrapCLPPerKg = 1.0
		}
	}
	return cogsSlice
}

func WriteCOGS(client *mongo.Client, dbName string, colName string, cogsSlice []model.COGS) {
	collection := ConnectToCollection(client, dbName, colName)

	models := make([]mongo.WriteModel, len(cogsSlice))
	for i, cogs := range cogsSlice {
		var model mongo.WriteModel
		iom := mongo.NewInsertOneModel()
		iom = iom.SetDocument(cogs)
		model = iom
		models[i] = model
	}
	res, err := collection.BulkWrite(context.TODO(), models)
	if err != nil {
		panic("WriteCOGS - BulkWrite Error: " + err.Error())
	}
	if res.InsertedCount != int64(len(cogsSlice)) {
		panic("Number of items written onto MongoDB is wrong")
	}
}

func CreateSalesValue() []model.SalesValue {
	salesValueSlice := make([]model.SalesValue, 68*3)
	for i := 0; i < 68; i++ {
		for j := 0; j < 3; j++ {
			salesValueSlice[i*3+j].ProductID = "T" + strconv.Itoa(i+1) + "_Product" + strconv.Itoa(j+1)
			salesValueSlice[i*3+j].CustomerID = "T" + strconv.Itoa(i+1) + "_Customer1"
			salesValueSlice[i*3+j].GrossSalesValuePerKg = 1.0
			salesValueSlice[i*3+j].NetSalesValuePerKg = 1.0
		}
	}
	return salesValueSlice
}

func WriteSalesValue(client *mongo.Client, dbName string, colName string, salesValueSlice []model.SalesValue) {
	collection := ConnectToCollection(client, dbName, colName)

	models := make([]mongo.WriteModel, len(salesValueSlice))
	for i, salesValue := range salesValueSlice {
		var model mongo.WriteModel
		iom := mongo.NewInsertOneModel()
		iom = iom.SetDocument(salesValue)
		model = iom
		models[i] = model
	}
	res, err := collection.BulkWrite(context.TODO(), models)
	if err != nil {
		panic("WriteSalesValue - BulkWrite Error: " + err.Error())
	}
	if res.InsertedCount != int64(len(salesValueSlice)) {
		panic("Number of items written onto MongoDB is wrong")
	}
}

func TestWriteCOGSAndSalesValue(t *testing.T) {
	client := ConnectToMongoDB()
	dbName := "automatic_tests"

	colName := "COGS"
	cogsSlice := CreateCOGS()
	WriteCOGS(client, dbName, colName, cogsSlice)

	colName = "SalesValue"
	salesValueSlice := CreateSalesValue()
	WriteSalesValue(client, dbName, colName, salesValueSlice)

	CloseMongoDBConnection(client)
}

func TestWriteCOGS(t *testing.T) {
	i := 15

	client := ConnectToMongoDB()
	dbName := "automatic_tests"
	colName := "COGS"

	cogsSlice := make([]model.COGS, 3)
	for j := 0; j < 3; j++ {
		cogsSlice[j].ProductID = "T" + strconv.Itoa(i) + "_Product" + strconv.Itoa(j+4)
		cogsSlice[j].COGSCLPPerKg = 1.0
		cogsSlice[j].ScrapCLPPerKg = 1.0
	}

	WriteCOGS(client, dbName, colName, cogsSlice)

	CloseMongoDBConnection(client)
}

func TestWriteSalesValue(t *testing.T) {
	i := 15
	j := 2

	client := ConnectToMongoDB()
	dbName := "automatic_tests"
	colName := "SalesValue"

	salesValueSlice := make([]model.SalesValue, 3)

	for k := 0; k < 3; k++ {
		salesValueSlice[k].ProductID = "T" + strconv.Itoa(i) + "_Product" + strconv.Itoa(k+4)
		salesValueSlice[k].CustomerID = "T" + strconv.Itoa(i) + "_Customer" + strconv.Itoa(j)
		salesValueSlice[k].GrossSalesValuePerKg = 1.0
		salesValueSlice[k].NetSalesValuePerKg = 1.0
	}

	WriteSalesValue(client, dbName, colName, salesValueSlice)

	CloseMongoDBConnection(client)
}
