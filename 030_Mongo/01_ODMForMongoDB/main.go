package main

import (
	"Go-Web-Dev/086_ODMForMongoDB/models"
	"fmt"

	"github.com/zebresel-com/mongodm"
)

func main() {
	// file, err := ioutil.ReadFile("locals.json")
	// if err != nil {
	// 	fmt.Printf("File error: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// var localMap map[string]map[string]string
	// json.Unmarshal(file, &localMap)

	// for key, values := range localMap {
	// 	fmt.Printf("Key: %v\n", key)
	// 	for k, v := range values {
	// 		fmt.Printf("Key: %v, Value: %v\n", k, v)
	// 	}
	// }

	dbConfig := &mongodm.Config{
		DatabaseHosts: []string{"localhost"},
		DatabaseName:  "mongodm_sample",
		// DatabaseUser:     "admin",
		// DatabasePassword: "admin",
		// The option `DatabaseSource` is the database used to establish
		// credentials and privileges with a MongoDB server. Defaults to the value
		// of `DatabaseName`, if that is set, or "admin" otherwise.
		// DatabaseSource: "admin",
		// Locals:         localMap["en-US"],
	}

	connection, err := mongodm.Connect(dbConfig)
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
	}

	connection.Register(&models.User{}, "users")
	connection.Register(&models.Message{}, "messages")
	connection.Register(&models.Customer{}, "customers")

	User := connection.Model("User")
	user := &models.User{}
	User.New(user) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	user.FirstName = "Max"
	user.LastName = "Mustermann"

	err = user.Save()
	if err != nil {
		fmt.Printf("Save new user error: %v\n", err)
	}

	msg := &models.Message{}
	MSG := connection.Model("Message")
	MSG.New(msg)
	msg.Sender = "Akagi"
	msg.Receiver = "Yushin"
	msg.Text = "Hello"
	err = msg.Save()
	if err != nil {
		fmt.Printf("Save new msg error: %v\n", err)
	}
}
