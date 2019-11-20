package main

import (
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/log"
)

func init() {
	log.SetSTDHook(5)
}

func main() {
	log.Info("Spinning up load balancer...")
	log.Info("Reading Config.yml...")
	proxy, err := config.ReadConfig()
	if err != nil {
		log.Error("An error occurred while trying to parse config.yml")
		log.Fatal(err)
	}

	// Test
	log.Info(proxy.URL())

	http.HandleFunc("/", proxy.Handler)
	log.Info("Listening to requests on port: " + proxy.Port)
	err = http.ListenAndServe(":"+proxy.Port, nil)
	if err != nil {
		log.Error("Failed to bind to port " + proxy.Port)
		log.Fatal("Make sure the port is available and not reserved")
	}
}
