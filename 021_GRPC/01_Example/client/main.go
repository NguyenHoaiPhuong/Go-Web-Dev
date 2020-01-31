package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/021_GRPC/01_Example/pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	checkError(err, "did not connect")
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Contact the server and print out its response
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	checkError(err, "could not greet")
	log.Printf("Greeting: %s", res.GetMessage())
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
