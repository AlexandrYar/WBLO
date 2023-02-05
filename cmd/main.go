package main

import (
	"github.com/AlexandrYar/WBLO/cmd/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routing.Run(router)
}
