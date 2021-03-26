package structs

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

type DBContext struct {
	echo.Context
	DB *sql.DB
}
