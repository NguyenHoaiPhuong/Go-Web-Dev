package services

import (
	"context"
	"log"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/04_Example/pb"
)

// UserSrvServer is used to implement pb.UserServer
type UserSrvServer struct {
	pb.UnimplementedUserSrvServer
}

// CreateUser : create new user
func (s *UserSrvServer) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	log.Println("Create new user")
	return &pb.CreateUserRes{}, nil
}

// ReadUser : read user info
func (s *UserSrvServer) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	log.Println("Read user info")
	return &pb.ReadUserRes{}, nil
}

// UpdateUser : update user with new info
func (s *UserSrvServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	log.Println("Update user info")
	return &pb.UpdateUserRes{}, nil
}

// DeleteUser : delete user
func (s *UserSrvServer) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	log.Println("Delete a specific user")
	return &pb.DeleteUserRes{}, nil
}
