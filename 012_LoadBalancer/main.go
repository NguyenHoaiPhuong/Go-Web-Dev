package main

import "github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/log"

func init() {
	log.SetSTDHook(5)
}

func main() {
	log.Info("App start ...")
}
