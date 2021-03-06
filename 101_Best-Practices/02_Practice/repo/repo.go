package repo

import (
	"Go-Web-Dev/101_Best-Practices/02_Practice/error"
	"Go-Web-Dev/101_Best-Practices/02_Practice/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoDB struct includes DB session
type MongoDB struct {
	Session *mgo.Session
}

// InitDBSession initalizes MongoDB session
func (db *MongoDB) InitDBSession(serverHost string) error.Error {
	var err error.Imp
	sess, osErr := mgo.Dial(serverHost)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorDBSessionInit)
		return err
	}
	db.Session = sess
	db.Session.SetMode(mgo.Monotonic, true)
	return nil
}

// EnsureIndex indexing
func (db *MongoDB) EnsureIndex(databaseName string, collectionName string, indexKey string) error.Error {
	var err error.Imp
	if db.Session == nil {
		err.InsertErrorMessage(error.ErrorDBSessionNil)
		err.InsertErrorMessage(error.ErrorDBIndex)
		return err
	}

	session := db.Session.Copy()
	index := mgo.Index{
		Key:        []string{indexKey},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	c := session.DB(databaseName).C(collectionName)
	osErr := c.EnsureIndex(index)
	if osErr != nil {
		err.InsertErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorDBIndex)
	}

	return nil
}

// GetAllDocuments gets all documents in the collection
// func (db *MongoDB) GetAllDocuments(databaseName string, collectionName string) ([]model.Book, error.Error) {
// 	var err error.Imp
// 	var books []model.Book

// 	session := db.Session.Copy()
// 	defer session.Close()

// 	c := session.DB(databaseName).C(collectionName)

// 	osErr := c.Find(bson.M{}).All(&books)
// 	if osErr != nil {
// 		err.SetErrorMessage(osErr.Error())
// 		err.InsertErrorMessage(error.ErrorAppGetAllBooks)
// 		return books, err
// 	}
// 	return books, nil
// }
func (db *MongoDB) GetAllDocuments(databaseName string, collectionName string) ([]*model.Book, error.Error) {
	var err error.Imp
	var docs []interface{}

	session := db.Session.Copy()
	defer session.Close()

	c := session.DB(databaseName).C(collectionName)

	osErr := c.Find(bson.M{}).All(&docs)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorAppGetAllBooks)
		return nil, err
	}

	books := make([]*model.Book, len(docs))
	for i, doc := range docs {
		book := new(model.Book)
		bsonBytes, _ := bson.Marshal(doc)
		bson.Unmarshal(bsonBytes, book)
		books[i] = book
	}
	return books, nil
}

// GetDocumentByKey gets document by given key and respective value
func (db *MongoDB) GetDocumentByKey(databaseName string, collectionName string, key string, value string) (model.Book, error.Error) {
	var err error.Imp
	var book model.Book

	session := db.Session.Copy()
	defer session.Close()

	c := session.DB(databaseName).C(collectionName)

	osErr := c.Find(bson.M{key: value}).One(&book)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorDBGetDocumentByKey)
		return book, err
	}
	return book, nil
}

// AddDocument adds new document
func (db *MongoDB) AddDocument(databaseName string, collectionName string, book *model.Book) error.Error {
	var err error.Imp

	session := db.Session.Copy()
	defer session.Close()

	c := session.DB(databaseName).C(collectionName)
	osErr := c.Insert(*book)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorDBDuplicatedKey)
		return err
	}
	return nil
}

// UpdateDocument update specific document
func (db *MongoDB) UpdateDocument(databaseName string, collectionName string, book *model.Book) error.Error {
	var err error.Imp

	session := db.Session.Copy()
	defer session.Close()

	c := session.DB(databaseName).C(collectionName)
	osErr := c.Update(bson.M{"isbn": book.ISBN}, book)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorDBUpdateDocument)
		return err
	}
	return nil
}

// DeleteDocumentByKey deletes document by given key and respective value
func (db *MongoDB) DeleteDocumentByKey(databaseName string, collectionName string, key string, value string) error.Error {
	var err error.Imp

	session := db.Session.Copy()
	defer session.Close()

	c := session.DB(databaseName).C(collectionName)

	osErr := c.Remove(bson.M{key: value})
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorDBGetDocumentByKey)
		return err
	}
	return nil
}
