package api

import (
	"math/rand"
	"net/http"
	"time"

	serializers "github.com/NguyenHoaiPhuong/Go-Web-Dev/000_test/wizeline/serialziers"
	"github.com/gin-gonic/gin"
)

// Routes :
func (s *Server) Routes() {
	s.g.GET("/", s.DefaultWelcome)
	api := s.g.Group("/v1")
	{
		api.GET("/", s.Welcome)
		api.GET("/shorten-url/:url", s.ShortenURL)
	}
}

// DefaultWelcome : ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Homepage")
}

// Welcome : ...
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, serializers.Resp{Result: "REST API Homepage"})
}

const (
	numOfLength = 10
)

var (
	urlToShortenUrl = make(map[string]string, 0)
	shortenUrlToUrl = make(map[string]string, 0)
)

// ShortenURL : ...
func (s *Server) ShortenURL(c *gin.Context) {
	url := c.DefaultQuery("url")
	var shortenURL string

	shortenURL, ok := urlToShortenUrl[url]
	if !ok {
		shortenURL = generateNewShortenUrl(url)
	}

	c.JSON(http.StatusOK, serializers.URLResp{URL: shortenURL})
}

func generateNewShortenUrl(url string) string {
	var (
		shortenUrl string
	)
	for true {
		str := RandStringRunes(numOfLength)
		if _, ok := shortenUrlToUrl[str]; !ok {
			shortenUrl = str
			urlToShortenUrl[url] = shortenUrl
			break
		}
	}
	return shortenUrl
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
