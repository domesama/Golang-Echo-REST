package routes

import (
	"fmt"
	"github.com/domesama/Golang-Echo-REST/handlers"
	"github.com/domesama/Golang-Echo-REST/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
)

//GetDefaultRoutes Serves listed routes
func GetDefaultRoutes(e *echo.Echo)  {
	//Default Example
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, fmt.Sprintf("%v", context.Request()))
	})

	//echo.HandlerFunc example
	e.GET("/products/:id", handlers.GetStuff, middlewares.CreateStuffMiddleware)
	e.GET("/waifus", handlers.GetAnimeWaifus)
	e.POST("/waifus", handlers.PostAnimeWaifu)
	e.DELETE("/waifus/:id", handlers.DeleteAnimeWaifu)
}
