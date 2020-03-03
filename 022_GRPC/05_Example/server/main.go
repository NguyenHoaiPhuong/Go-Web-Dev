package main

import (
	"context"
	"log"
	"net"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/05_Example/errors"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/05_Example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	err := errors.New("backend", int32(codes.PermissionDenied), "Say Hello denied")
	err = errors.WithMessage(err, "server")
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	errors.CheckError(err, "failed to listen")
	svr := grpc.NewServer()
	pb.RegisterGreeterServer(svr, &server{})
	err = svr.Serve(lis)
	errors.CheckError(err, "failed to serve")
}
