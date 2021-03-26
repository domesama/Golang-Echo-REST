package middlewares

import (
	"github.com/domesama/Golang-Echo-REST/packages/structs"
	"github.com/labstack/echo/v4"
)

func CreateStuffMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		p := structs.NewStuff([]string{"OwO", "UwU", "Anime", "Waifus"})
		productCtx := &structs.StuffContext{Context: c, Stuff: p }
		return next(productCtx)
	}
}
