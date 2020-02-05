package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/10_Example/pb/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("user:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	request := &pb.FindRequest{
		Id: "123456",
	}
	resp, err := client.FindUser(context.Background(), request)
	// To do something with resp from instance server response
	fmt.Println(resp)
}
