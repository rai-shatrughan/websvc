package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/router/gin"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "DEBUG", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")

	pwd, _ := os.Getwd()
	absConfPath, _ := filepath.Abs(pwd + "/gateway/conf/config.json")
	configFile := flag.String("c", absConfPath, "Path to the configuration filename")

	flag.Parse()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[LURA]")

	routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	log.Printf("Starting Lura Gateway at :" + strconv.Itoa(serviceConfig.Port))
	routerFactory.New().Run(serviceConfig)
}
