package main

import (
	"context"
	"log"
	"net"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/02_GRPC/01_Example/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement pb.GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	checkError(err, "failed to listen")
	svr := grpc.NewServer()
	pb.RegisterGreeterServer(svr, &server{})
	err = svr.Serve(lis)
	checkError(err, "failed to serve")
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
