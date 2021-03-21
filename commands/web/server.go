package web

import (
	"fmt"
	"github.com/domesama/Golang-Echo-REST/routes"
	"github.com/domesama/Golang-Echo-REST/validators"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"os"
)

var e = echo.New()

//Start starts the application
func Start()  {

	port := os.Getenv("DEV_PORT")
	if port == ""{
		port = "8080"
	}

	//Add Custom Validator
	e.Validator = &validators.CustomValidator{ Validator: validator.New()}

	//Routes
	routes.GetDefaultRoutes(e)

	e.Logger.Print(fmt.Sprintf("Started server on http://localhost:%s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}




