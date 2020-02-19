package utils

// FirstName string, First Name
var FirstName = "Nguyen"

var middleName = "Hoai"

// LastName string, Last Name
var LastName = "Phuong"

// FullName return a string of First + Middle + Last name
func FullName() string {
	return FirstName + " " + middleName + " " + LastName
}
