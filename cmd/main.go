package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"rest_api_test/internal/handler"
	"rest_api_test/internal/repository"
	"rest_api_test/internal/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load("../.env"); err != nil { //"../.env" for local using
		logrus.Fatalf("error loading env variables :%s", err.Error())
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("error while connectig to db :%s", err.Error())
	}

	repos := repository.NewRepository(db)

	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	//e.GET("/", handler.Hello)
	e.GET("/", handlers.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
