package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	nltk "github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/02_Example/golang/nltk_service"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client nltk.KeywordServiceClient
}

const SERVER_ADDR = "127.0.0.1:6000"

func InitGrpcConnection() (*GrpcClient, error) {
	conn, err := grpc.Dial(SERVER_ADDR, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := nltk.NewKeywordServiceClient(conn)
	return &GrpcClient{conn, client}, nil
}

func (g *GrpcClient) MyKeywords(text string) ([]string, error) {
	req := nltk.Request{
		Text: text,
	}

	res, err := g.client.GetKeywords(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	return res.Keywords, nil
}

func main() {
	client, err := InitGrpcConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter some text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println("Keywords:")
		keywords, err := client.MyKeywords(text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(keywords)
		fmt.Println()
	}
}
