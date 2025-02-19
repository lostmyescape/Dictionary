package main

import (
	"dictionary/internal/service"
	"dictionary/pkg/logs"
	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// api words
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/word", svc.UpdateWords)
	api.DELETE("/words/:id", svc.DeleteWords)
	// api search
	api.GET("/search/ru", svc.SearchWords)

	// api reports
	api.GET("/report/:id", svc.GetReport)
	api.POST("/reports", svc.CreateReport)
	api.PUT("/report/:id", svc.UpdateReport)
	api.DELETE("/report/:id", svc.DeleteReport)

	router.Logger.Fatal(router.Start(":8000"))
}
