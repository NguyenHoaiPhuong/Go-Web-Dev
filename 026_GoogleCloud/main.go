package main

import (
	"fmt"
	"log"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/026_GoogleCloud/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/026_GoogleCloud/storage"
)

func main() {
	conf := config.GetConfig()
	gcs := storage.Init(conf)
	client, err := gcs.InitClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List all objests in the bucket")
	objs, err := storage.List(client, conf.GoogleBucketName)
	if err != nil {
		log.Fatal(err)
	}
	for _, obj := range objs {
		fmt.Println(obj)
	}

	// fmt.Println("List all objests with prefix in the bucket")
	// objs, err := storage.ListByPrefix(client, conf.GoogleBucketName, "images/integration-test", "/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, obj := range objs {
	// 	fmt.Println(obj)
	// 	storage.Delete(client, conf.GoogleBucketName, obj)
	// }
}
