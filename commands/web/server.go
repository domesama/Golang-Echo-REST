package web

import (
	"fmt"
	"github.com/domesama/Golang-Echo-REST/configs"
	"github.com/domesama/Golang-Echo-REST/middlewares"
	"github.com/domesama/Golang-Echo-REST/routes"
	"github.com/domesama/Golang-Echo-REST/validators"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var appConf = &configs.ConfigApp{}

func init()  {
	err := cleanenv.ReadEnv(appConf)
	if err != nil{
		e.Logger.Fatal(err)
	}
	fmt.Println( "The current app configs are as follows:\n" ,fmt.Sprintf("%v", appConf))
}

//Start starts the application
func Start()  {

	//Deprecated -> Go use some env lib for prod uses
	//port := os.Getenv("DEV_PORT")

	//Add Custom Validator
	e.Validator = &validators.CustomValidator{ Validator: validator.New()}
	e.Use(middlewares.ProductMiddleware)
	//Routes
	routes.GetDefaultRoutes(e)

	//e.Logger.Print(fmt.Sprintf("Started server on http://localhost:%s", appConf.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", appConf.Host,appConf.Port)))
}




