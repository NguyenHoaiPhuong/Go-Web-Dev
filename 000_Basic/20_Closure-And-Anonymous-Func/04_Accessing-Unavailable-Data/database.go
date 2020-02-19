package main

// Database structure
type Database struct {
	URL string
}

// NewDatabasae returns struct Database
func NewDatabasae(url string) Database {
	return Database{URL: url}
}
