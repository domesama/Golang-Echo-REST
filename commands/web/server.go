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

//Init Echo Framework & add AppConfig pointers
var e = echo.New()
var appConf = &configs.ConfigApp{}

func init()  {
	//Reads the given env according to that pre-defined config
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

	//Add global middleware, note middlewares could be added later on for EACH individual function
	e.Use(middlewares.UseDatabaseMiddleWare)

	//Routes
	routes.GetDefaultRoutes(e)

	//Starts the server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", appConf.Host,appConf.Port)))
}




