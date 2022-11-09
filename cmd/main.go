package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"rest_api_test/internal/handler"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load("../.env"); err != nil { //"../.env" for local using
		logrus.Fatalf("error loading env variables :%s", err.Error())
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
