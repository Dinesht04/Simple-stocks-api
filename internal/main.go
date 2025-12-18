package main

import (
	"github.com/dinesht04/basic-stock-api/internal/router"
)

func main() {
	r := router.CreateRouter()

	r.Run("localhost:3000")
}
