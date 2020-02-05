package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/03_Example/pb"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

// server is used to implement pb.ArithmeticServer
type server struct {
	pb.UnimplementedArithmeticServer
}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a := in.GetA()
	b := in.GetB()
	res := a + b
	log.Printf("%v + %v = %v\n", a, b, res)
	return &pb.Response{Result: res}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a := in.GetA()
	b := in.GetB()
	res := a - b
	log.Printf("%v - %v = %v\n", a, b, res)
	return &pb.Response{Result: res}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a := in.GetA()
	b := in.GetB()
	res := a * b
	log.Printf("%v x %v = %v\n", a, b, res)
	return &pb.Response{Result: res}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a := in.GetA()
	b := in.GetB()
	if b == 0 {
		err := errors.New("Divided by 0")
		log.Println("Divided by 0")
		return &pb.Response{}, err
	}
	res := a / b
	log.Printf("%v : %v = %v\n", a, b, res)
	return &pb.Response{Result: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	checkError(err, "failed to listen")
	svr := grpc.NewServer()
	pb.RegisterArithmeticServer(svr, &server{})
	err = svr.Serve(lis)
	checkError(err, "failed to serve")
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
