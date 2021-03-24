package structs

import "github.com/labstack/echo/v4"

type Product []map[int]string

type ProductContext struct {
	echo.Context
	Product
}

func NewProduct(products []string) Product{
	prod := Product{}
	for index,val := range products{
		prod = append(prod, map[int]string{index:val})
	}
	return prod
}