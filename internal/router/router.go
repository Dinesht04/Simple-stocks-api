package router

import (
	"net/http"

	"github.com/dinesht04/basic-stock-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func RateLimitingMiddleWare() {

}

func CreateRouter(db *xorm.Engine) *gin.Engine {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/add", middleware.InsertStock(db))
	router.POST("/sell", middleware.RemoveStock(db))
	router.GET("/", middleware.ListStocks(db))

	return router
}
