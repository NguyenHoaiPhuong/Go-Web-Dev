package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/03_Example/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"

	defaultA = 1.0
	defaultB = 2.0
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	checkError(err, "did not connect")
	defer conn.Close()
	client := pb.NewArithmeticClient(conn)

	a := defaultA
	b := defaultB
	if len(os.Args) > 2 {
		a, err = strconv.ParseFloat(os.Args[1], 64)
		checkError(err, "could not parse float")
		b, err = strconv.ParseFloat(os.Args[2], 64)
		checkError(err, "could not parse float")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Add(ctx, &pb.Request{A: a, B: b})
	checkError(err, "could not do addition")
	log.Printf("Addition: %f", res.GetResult())

	res, err = client.Subtract(ctx, &pb.Request{A: a, B: b})
	checkError(err, "could not do subtraction")
	log.Printf("Subtraction: %f", res.GetResult())

	res, err = client.Multiply(ctx, &pb.Request{A: a, B: b})
	checkError(err, "could not do multiplication")
	log.Printf("Multiplication: %f", res.GetResult())

	res, err = client.Divide(ctx, &pb.Request{A: a, B: b})
	checkError(err, "could not do division")
	log.Printf("Division: %f", res.GetResult())
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
