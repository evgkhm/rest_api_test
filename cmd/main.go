package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"rest_api_test/internal/handler"
	"rest_api_test/internal/repository"
	"rest_api_test/internal/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load("./.env"); err != nil { //"../.env" for local using
		logrus.Fatalf("error loading env variables :%s", err.Error())
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("error while connectig to db :%s", err.Error())
	}

	repos := repository.NewRepository(db)

	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	r := handlers.InitRoutes()
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, handlers)
	})*/
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
