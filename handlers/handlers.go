package handlers

import (
	"fmt"
	"github.com/domesama/Golang-Echo-REST/packages/structs"
	"github.com/domesama/Golang-Echo-REST/services"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func GetStuff(ctx echo.Context)  error{
	prodCtx := ctx.(*structs.StuffContext)
	param := ctx.Param("id")

	if param == ""{
		ctx.Logger().Print("The required parameter was not given")
	}
	id,err := strconv.Atoi(param)

	if err != nil{
		ctx.Logger().Print("There was an error parsing the given params")
	}

	for _,val := range prodCtx.Stuff{
		for index,_ := range val{
			if id == index{
				ctx.JSON(http.StatusOK, val)
				return nil
			}
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "There was no Products within the given id")
}

func GetAnimeWaifus(ctx echo.Context) (err error){
	db := ctx.(*structs.DBContext)
	defer db.DB.Close()
	svc := services.NewAnimeWaifuService(db.DB)

	waifu,err := svc.GetWaifus()
	if err != nil || &waifu == nil{
		log.Fatal(err)
		return err
	}
	return ctx.JSON(http.StatusOK, waifu)
}

func PostAnimeWaifu(ctx echo.Context) (err error) {
	db := ctx.(*structs.DBContext)
	defer db.DB.Close()
	svc := services.NewAnimeWaifuService(db.DB)

	waifu := structs.AnimeWaifu{}

	if err = ctx.Bind(&waifu); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(waifu); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = svc.SendWaifus(waifu); err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK,waifu)
}

func DeleteAnimeWaifu(ctx echo.Context) (err error){
	db := ctx.(*structs.DBContext)
	defer db.DB.Close()
	svc := services.NewAnimeWaifuService(db.DB)
	id,_ := strconv.Atoi(ctx.Param("id"))

	err = svc.DeleteWaifu(id)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, fmt.Sprintf(`Your waifu with id "%d" was deleted ): very sad`, id ))
}