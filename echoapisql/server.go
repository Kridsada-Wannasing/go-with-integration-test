package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	echo "github.com/labstack/echo/v4"

	"github.com/Kridsada-Wannasing/echoapisql/user"
)

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func main() {
	user.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "apidesign" && password == "45678" {
			return true, nil
		}

		return false, nil
	}))

	e.GET("/users", user.GetUsersHandler)
	e.GET("/users/:id", user.GetUserHandler)
	e.POST("/users", user.CreateUsersHandler)
	e.PATCH("/users/:id", user.UpdateUsersHandler)
	e.DELETE("/users/:id", user.DeleteUserByID)

	log.Println("Server started at :2565")
	log.Fatal((e.Start(":2565")))
	log.Println("bye bye!")
}
