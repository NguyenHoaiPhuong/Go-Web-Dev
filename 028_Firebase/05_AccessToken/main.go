package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/028_Firebase/05_AccessToken/config"
	"golang.org/x/oauth2/google"
)

func main() {
	ctx := context.Background()

	conf := config.GetConfig()
	data, err := ioutil.ReadFile(conf.GoogleCredentials)
	if err != nil {
		log.Fatal(err)
	}
	creds, err := google.CredentialsFromJSON(ctx, data, "https://www.googleapis.com/auth/bigquery")
	if err != nil {
		log.Fatal(err)
	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token.AccessToken)

}
