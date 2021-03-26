package middlewares

import (
	"github.com/domesama/Golang-Echo-REST/configs"
	"github.com/domesama/Golang-Echo-REST/packages/structs"
	"github.com/labstack/echo/v4"
	"log"
)

func UseDatabaseMiddleWare(next echo.HandlerFunc)  echo.HandlerFunc{
	return func(e echo.Context) error{
		db, err :=configs.InitDB()
		if err != nil{
			log.Fatal(err)
		}
		dbContext := &structs.DBContext{Context: e,DB: db}
		return next(dbContext)
	}
}
