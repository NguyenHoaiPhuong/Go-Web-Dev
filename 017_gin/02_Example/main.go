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
	person := parsePerson(c)

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
	id, found := validatePersonID(c)
	if found {
		// oldPerson := persons[id]
		newPerson := parsePerson(c)
		newPerson.ID = id
		persons[id] = newPerson

		c.JSON(200, gin.H{
			"ID":   newPerson.ID,
			"Name": newPerson.Name,
			"Age":  newPerson.Age,
		})
	} else {
		c.String(http.StatusOK, "No person matches the ID %v", id)
	}
}

func deletePerson(c *gin.Context) {
	id, found := validatePersonID(c)
	if found {
		person := persons[id]
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

func validatePersonID(c *gin.Context) (id int, found bool) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	checkError(err)

	if _, ok := persons[id]; !ok {
		return id, false
	}

	return id, true
}

func parsePerson(c *gin.Context) *Person {
	person := new(Person)
	err := c.Bind(person)
	checkError(err)

	return person
}
