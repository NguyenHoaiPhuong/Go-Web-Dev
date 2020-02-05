package main

import (
	"log"
	"net"

	pb "github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/10_Example/pb/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	request []*pb.GetRequest
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *server) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}

func (s *server) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServer(srv, &server{})
	srv.Serve(lis)
}
