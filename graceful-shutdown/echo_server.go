package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK)
	})

	go func() {
		if err := e.Start(":1234"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
