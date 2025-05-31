package main

import (
	"github.com/labstack/echo/v4"

	"net/http"
	"time"
)

func main() {
	e := echo.New()

	delaySeconds := 3	

	e.GET(`/:text`, func(c echo.Context) error {
			param := c.Param("text")

			time.Sleep(time.Duration(delaySeconds) * time.Second)

			return c.String(http.StatusOK, param)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
