Golang RESTful API using GORM and Gorilla Mux

REST:
- REST stands for Representational State Transfer
- It is a very streamlined and lightweight web service compared to SOAP or WSDL
- Performance, scalability, simplicity, portability are the core principles behind the REST API

RESTful Web APIs:
- The essence of the REST architecture comprises a client and a server
- The REST API allows diverse systems to connect and send/receive data in the direct way
- The server accepts the incoming messages, then replies to it, while the client creates the connection, then delivers messages to the server
- The RESTful client would be an HTTP client, and the RESTful server would be the HTTP server
- Each and every REST API call has a relationship between an HTTP verb and the URL
- The reserves (data or business logic) in the database in an application can be defined with an API endpoint in the REST.

Installation:
- go get -u github.com/gorilla/mux
- go get -u github.com/jinzhu/gorm
- go get -u github.com/go-sql-driver/mysql