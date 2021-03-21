package handlers

import (
	"github.com/domesama/Golang-Echo-REST/packages/structs"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)



func GetProducts(ctx echo.Context)  error{

	prodCtx := ctx.(*structs.ProductContext)

	param := ctx.Param("id")
	if param == ""{
		ctx.Logger().Print("The required parameter was not given")
	}
	id,err := strconv.Atoi(param)

	if err != nil{
		ctx.Logger().Print("There was an error parsing the given params")
	}

	for _,val := range prodCtx.Prod{
		for index,_ := range val{
			if id == index+1{
				ctx.JSON(http.StatusOK, val)
			}
		}
	}
	return nil
}

func GetAnimeWaifus(ctx echo.Context) (err error) {

	defer ctx.Request().Body.Close()

	waifu := structs.AnimeWaifu{}

	if err = ctx.Bind(&waifu); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(waifu); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK,waifu)
}