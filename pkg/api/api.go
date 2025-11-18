package api

import (
    "github.com/gin-gonic/gin"
    "github.com/varrahan/url-shortener/internal/api/handler"
    "github.com/varrahan/url-shortener/internal/api/store"
)

func Init() {
    store.InitStore()
}

func RegisterRoutes(r *gin.Engine) {
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the URL Shortener API",
        })
    })

    r.POST("/create-short-url", handler.CreateShortUrl)
    r.GET("/:shortUrl", handler.HandleShortUrlRedirect)
}
