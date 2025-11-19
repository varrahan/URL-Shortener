package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/varrahan/url-shortener/internal/api/utils"
	"github.com/varrahan/url-shortener/pkg/api"
)

func main() {
	utils.LoadEnv(".env")

	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	api.Init()

	api.RegisterRoutes(r)

	port := utils.GetEnv("INTERNAL_PORT", "9000")

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
