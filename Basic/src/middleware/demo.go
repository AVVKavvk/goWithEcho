package custommiddleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc{
	return func (c echo.Context) error  {
		fmt.Println("Inside middleware")
		// return echo.ErrBadGateway;
		return next(c)
	}
}