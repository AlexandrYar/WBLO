package routing

import (
	handlers "github.com/AlexandrYar/WBLO/internal"
	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.LoadHTMLGlob("tmp/html/*.html")
	router.GET("/", internal.Index)
	router.POST("/", internal.Index)
	router.GET("/json", internal.Json)
	router.POST("/json", internal.Json)
	handlers.Index()
	router.Run("localhost:8080")
}
