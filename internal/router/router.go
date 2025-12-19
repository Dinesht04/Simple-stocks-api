package router

import (
	"net/http"

	"github.com/dinesht04/basic-stock-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func DBMiddleware(db *xorm.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func CreateRouter(db *xorm.Engine) *gin.Engine {

	router := gin.Default()
	router.Use(DBMiddleware(db))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/add", middleware.InsertStock(db))

	return router
}
