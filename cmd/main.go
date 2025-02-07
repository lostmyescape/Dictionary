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

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)

	api.POST("/words", svc.CreateWords)

	api.PUT("/words/id", svc.UpdateWords)

	api.DELETE("/words/:id", svc.DeleteWords)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))

	//api.POST("/words", func(c echo.Context) error {
	//	var w word
	//	if err := c.Bind(&w); err != nil {
	//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	//	}
	//
	//	// Пример сохранения в базу данных (если настроено подключение)
	//	_, err := db.Exec("INSERT INTO words (word) VALUES ($1)", w.Word)
	//	if err != nil {
	//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert data"})
	//	}
	//
	//	return c.JSON(http.StatusCreated, w)
	//})

}
