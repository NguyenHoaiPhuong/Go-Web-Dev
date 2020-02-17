// Using GET, POST, PUT, PATCH, DELETE and OPTIONS

package main

import (
	"errors"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Person struct
type Person struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"Name" bson:"Name"`
	Age  int    `json: "Age" bson: "Age"`
}

// Persons : map of person and its ID
type Persons map[int]*Person

var persons Persons

func init() {
	persons = make(Persons, 0)
}

func main() {
	router := gin.Default()
	router.GET("/persons", getAllPersons)
	router.POST("/person", addNewPerson)
	router.PUT("/person/:id", modifyPerson)
	router.DELETE("/person/:id", deletePerson)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)
	router.Run(":9000")
}

func getAllPersons(c *gin.Context) {
	out, err := json.Marshal(persons)
	checkError(err)

	c.String(http.StatusOK, "List of persons: %v", string(out))
}

func addNewPerson(c *gin.Context) {
	person := new(Person)
	err := c.Bind(person)
	checkError(err)

	id := person.ID
	if _, ok := persons[id]; ok {
		err := errors.New("Person ID " + strconv.Itoa(person.ID) + " exists.")
		checkError(err)
		return
	}
	persons[id] = person

	c.JSON(200, gin.H{
		"ID":   person.ID,
		"Name": person.Name,
		"Age":  person.Age,
	})
}

func modifyPerson(c *gin.Context) {

}

func deletePerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	checkError(err)

	if person, ok := persons[id]; ok {
		delete(persons, id)
		c.JSON(200, gin.H{
			"ID":   person.ID,
			"Name": person.Name,
			"Age":  person.Age,
		})
	} else {
		c.String(http.StatusOK, "No person matches the ID %v", id)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
