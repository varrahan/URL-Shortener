package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/varrahan/url-shortener/internal/api/utils"
	"github.com/varrahan/url-shortener/pkg/api"
)

func main() {
	utils.LoadEnv(".env")

	r := gin.Default()

	api.Init()

	api.RegisterRoutes(r)

	port := utils.GetEnv("INTERNAL_PORT", "9000")
	env := utils.GetAllEnv(".env")

	log.Printf("%v", env)

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
