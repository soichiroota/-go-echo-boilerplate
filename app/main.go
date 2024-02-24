package main

import (
	"sample/db"
	"sample/models"
	"sample/routes"

	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  db.Init()

	err := models.NewCreateUserTable()
	if err != nil {
		log.Fatal(err)
	}
	err = models.NewCreateTodoTable()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(routes.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}