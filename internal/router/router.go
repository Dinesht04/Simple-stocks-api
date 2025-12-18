package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
