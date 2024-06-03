package main

import (
	"github.com/09184417478m/Basket_Midterm/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user/login", handlers.Login)
	e.GET("/basket/", handlers.GetBaskets)
	e.POST("/basket/", handlers.CreateBasket)
	e.PATCH("/basket/:id", handlers.UpdateBasket)
	e.GET("/basket/:id", handlers.GetBasket)
	e.DELETE("/basket/:id", handlers.DeleteBasket)

	e.Start(":8080")
}
