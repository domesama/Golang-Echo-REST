package structs

import (
	"github.com/labstack/echo/v4"
)

type Stuff []map[int]string

type StuffContext struct {
	echo.Context
	Stuff
}

func NewStuff(s []string) Stuff {
	stuff := Stuff{}
	for index,val := range s{
		stuff = append(stuff, map[int]string{index:val})
	}
	return stuff
}