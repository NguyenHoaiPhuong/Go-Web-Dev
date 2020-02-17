// Using GET, POST, PUT, PATCH, DELETE and OPTIONS

package main

import (
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Person struct
type Person struct {
	Name string `json:"Name" bson:"Name"`
	Age  int    `json: "Age" bson: "Age"`
}

// Persons : slice of persons
type Persons []*Person

var persons Persons

func init() {
	persons = make(Persons, 0)
}

func main() {

	r := gin.Default()
	r.GET("someGet", getting)
	r.POST("somePost", posting)
	r.Run(":9000")
}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"verb": "GET",
	})
}

func posting(c *gin.Context) {
	person := new(Person)
	err := c.Bind(person)
	checkError(err)
	persons = append(persons, person)

	c.JSON(200, gin.H{
		"verb":       "POST",
		"personName": person.Name,
		"personAge":  person.Age,
	})

	out, err := json.Marshal(persons)
	checkError(err)

	c.String(http.StatusOK, "List of persons: %v", string(out))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
