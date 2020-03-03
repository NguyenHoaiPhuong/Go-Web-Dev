package api

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/04_Example/client/models"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/04_Example/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Homepage : "/"
func Homepage(c *gin.Context) {
	c.JSON(http.StatusOK, "Homepage")
}

// RestAPIHomepage : "/api"
func RestAPIHomepage(c *gin.Context) {
	c.JSON(http.StatusOK, "REST API Homepage")
}

// CreateUser : "/api/user/register"
func CreateUser(c *gin.Context) {
	log.Println("Create new user ...")

	userInfo := new(models.User)
	err := c.ShouldBind(userInfo)
	checkError(err, "cannot parse user info from request body")

	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure(), grpc.WithBlock())
	checkError(err, "did not connect")
	defer conn.Close()
	client := pb.NewUserSrvClient(conn)
	client.CreateUser(context.Background(), &pb.CreateUserReq{
		User: &pb.UserInfo{
			Id:       uint32(userInfo.ID),
			FullName: userInfo.FullName,
			Username: userInfo.UserName,
			Email:    userInfo.Email,
			Password: userInfo.Password,
		},
	})
}

// ReadUserProfile : "/api/user/profile/:id"
func ReadUserProfile(c *gin.Context) {
	log.Println("Read user profile ...")

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	checkError(err, "cannot convert string id to integer id")

	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure(), grpc.WithBlock())
	checkError(err, "did not connect")
	defer conn.Close()
	client := pb.NewUserSrvClient(conn)
	client.ReadUser(context.Background(), &pb.ReadUserReq{
		User: &pb.UserInfo{
			Id: uint32(id),
		},
	})
}

// UpdateUserProfile : "/api/user/profile"
func UpdateUserProfile(c *gin.Context) {
	log.Println("Update user profile ...")

}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
