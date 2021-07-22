package main

import (
	"flag"
	"log"

	"src"

	"github.com/gin-gonic/gin"
)

func main() {
	// global configuration
	log.SetFlags(log.Ldate | log.Ltime)

	// handle command line flags
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to an optional config file")
	flag.Parse()

	// load configuration file if it exists
	var config *src.ApiConfig
	if configPath != "" {
		conf, err := src.NewApiConfigFromFile(configPath)
		if err != nil {
			log.Fatal(err)
		}
		config = conf
	}

	// connect to SingleStore
	db, err := src.NewSingleStore(config.SingleStore)
	if err != nil {
		log.Fatal(err)
	}

	// we will use gin as our http server and router
	router := gin.Default()

	// register the api
	api := src.NewApi(db)
	api.RegisterRoutes(router)

	router.Run(":8000")
}
