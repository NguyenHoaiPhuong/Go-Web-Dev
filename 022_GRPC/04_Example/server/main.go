package main

import (
	"log"
	"net"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/04_Example/pb"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/04_Example/server/services"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

func main() {
	lis, err := net.Listen("tcp", port)
	checkError(err, "failed to listen")
	svr := grpc.NewServer()
	pb.RegisterUserSrvServer(svr, &services.UserSrvServer{})
	err = svr.Serve(lis)
	checkError(err, "failed to serve")
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
