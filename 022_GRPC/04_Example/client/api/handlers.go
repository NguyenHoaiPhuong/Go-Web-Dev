package api

import (
	"log"
	"net/http"

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
	userInfo := new(models.User)
	err := c.ShouldBind(userInfo)
	checkError(err, "cannot parse user info from request body")

	// Set up a connection to the server
	srvAddr := "localhost:9001"
	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure(), grpc.WithBlock())
	checkError(err, "did not connect")
	defer conn.Close()
	client := pb.NewUserSrvClient(conn)
	client.CreateUser(c.Request.Context(), &pb.CreateUserReq{
		User: &pb.UserInfo{
			Id:       uint32(userInfo.ID),
			FullName: userInfo.FullName,
			Username: userInfo.UserName,
			Email:    userInfo.Email,
			Password: userInfo.Password,
		},
	})
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+":", err)
	}
}
