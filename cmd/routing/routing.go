package routing

import (
	"github.com/AlexandrYar/WBLO/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.LoadHTMLGlob("tmp/html/*.html")
	router.GET("/", handlers.Index)
	router.POST("/", handlers.Index)
	router.GET("/json", handlers.Json)
	router.POST("/json", handlers.Json)

	router.Run("localhost:8080")
}
