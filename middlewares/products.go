package middlewares

import (
	"github.com/domesama/Golang-Echo-REST/packages/structs"
	"github.com/labstack/echo/v4"
)

func ProductMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		p := structs.NewProduct([]string{"OwO", "UwU", "Anime", "Waifus"})
		productCtx := &structs.ProductContext{Context: c, Products: p }
		return next(productCtx)
	}
}
