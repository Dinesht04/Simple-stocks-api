package data

import (
	"github.com/dinesht04/basic-stock-api/internal/middleware"
	"github.com/dinesht04/basic-stock-api/internal/model"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func ConnectToDB() *xorm.Engine {
	engine, err := xorm.NewEngine("postgres", "user=postgres password=password host=localhost port=5432 dbname=stocks sslmode=disable")
	middleware.Check(err)

	err = engine.DB().Ping()
	middleware.Check(err)

	err = engine.Sync(
		new(model.Stock),
	)
	middleware.Check(err)
	return engine
}
