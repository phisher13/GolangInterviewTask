package main

import (
	"log"
	"urls/config"
	"urls/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("../config")
	
	if err != nil {
		log.Panic("Config isn't loaded")
	} else {
		router := gin.Default()
		public := router.Group("/")
		routes.Router(public)


		router.Run(cfg.Server.Port)
	}
}