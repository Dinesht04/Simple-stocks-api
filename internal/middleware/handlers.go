package middleware

import (
	"fmt"
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
		ExistingStock := new(model.Stock)

		has, err := db.Where("stock_id = ?", Stock.StockId).Get(ExistingStock)
		if has == false {
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
				"message": "Stock bought Successfully",
			})

		} else {
			fmt.Println(ExistingStock)
			_, err = db.Where("stock_id = ?", Stock.StockId).Update(&model.Stock{
				Quantity: Stock.Quantity + ExistingStock.Quantity,
			})
			ctx.JSON(http.StatusOK, gin.H{
				"message":            "Stock exists already, added buy order",
				"prev_quantity":      ExistingStock.Quantity,
				"new_total_quantity": Stock.Quantity + ExistingStock.Quantity,
			})

		}

	}

}

func RemoveStock(db *xorm.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Details DeleteRequest
		err := ctx.BindJSON(&Details)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Format",
			})
		}

		Stock := new(model.Stock)

		has, err := db.Where("stock_id = ?", Details.StockId).Get(Stock)
		if has == false {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Stock doesnt exist",
			})
		}

		if Stock.Quantity-Details.Quantity < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Not enough stocks in inventory",
			})
		}

		Stock.Quantity = Stock.Quantity - Details.Quantity
		_, err = db.Where("u_u_i_d = ?", Stock.UUID).Update(&model.Stock{Quantity: Stock.Quantity})
		Check(err)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Stock sold successfully",
		})

	}
}
