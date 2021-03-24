package routes

import (
	"fmt"
	"github.com/domesama/Golang-Echo-REST/handlers"
	"github.com/domesama/Golang-Echo-REST/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Serves listed routes
func GetDefaultRoutes(e *echo.Echo)  {

	//Default Example
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, fmt.Sprintf("%v", context.Request().Close))
	})

	//echo.HandlerFunc example
	e.POST("/waifus", handlers.GetAnimeWaifus)
	e.GET("/products/:id", handlers.GetProducts, middlewares.ProductMiddleware)
}
