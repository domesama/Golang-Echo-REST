package web

import (
	"fmt"
	"github.com/domesama/Golang-Echo-REST/configs"
	"github.com/domesama/Golang-Echo-REST/routes"
	"github.com/domesama/Golang-Echo-REST/validators"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var config = &configs.ConfigDatabase{}

func init()  {
	err := cleanenv.ReadEnv(config)
	if err != nil{
		e.Logger.Fatal(err)
	}
	fmt.Println(config)
}

//Start starts the application
func Start()  {
	//port := os.Getenv("DEV_PORT")


	//Add Custom Validator
	e.Validator = &validators.CustomValidator{ Validator: validator.New()}

	//Routes
	routes.GetDefaultRoutes(e)

	e.Logger.Print(fmt.Sprintf("Started server on http://localhost:%s", config.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", config.Port)))
}




