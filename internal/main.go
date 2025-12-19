package main

import (
	"github.com/dinesht04/basic-stock-api/internal/data"
	"github.com/dinesht04/basic-stock-api/internal/router"
)

func main() {
	db := data.ConnectToDB()
	r := router.CreateRouter(db)

	r.Run("localhost:3000")
}
