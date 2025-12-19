package middleware

import (
	"net/http"

	"github.com/dinesht04/basic-stock-api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"xorm.io/xorm"
)

type CreateRequest struct {
	StockId  string `json:"stock_id" validate:"required"`
	Name     string `json:"string" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
	Company  string `json:"company" validate:"required"`
}

type DeleteRequest struct {
	StockId  string `json:"stock_id" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
}

func InsertStock(db *xorm.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Stock CreateRequest
		err := ctx.BindJSON(&Stock)
		Check(err)

		InseringStock := &model.Stock{
			UUID:     uuid.NewString(),
			StockId:  Stock.StockId,
			Name:     Stock.Name,
			Price:    Stock.Price,
			Quantity: Stock.Quantity,
			Company:  Stock.Company,
		}

		_, err = db.Insert(InseringStock)
		Check(err)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	}

}

func RemoveStock(db *xorm.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Details DeleteRequest
		err := ctx.BindJSON(&Details)
		Check(err)

		var Stock model.Stock
		has, err := db.Where("stock = ?", Details.StockId).Exist(&Stock)
		if has == false {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Stock doesnt exist",
			})
		}

	}
}
