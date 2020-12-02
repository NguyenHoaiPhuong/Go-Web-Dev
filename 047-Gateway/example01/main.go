package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router/gin"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	flag.Parse()
	parser := config.NewParser()

	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Fatal("ERROR:", err.Error())
	}
	fmt.Println("2222")
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}
	fmt.Printf("%+v\n", serviceConfig)

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")

	routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	routerFactory.New().Run(serviceConfig)
}
