package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// GetReport ищем репорт по айди
func (s *Service) GetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	report, err := repo.GetReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: report})
}

// CreateReport Добавляем новый репорт в БД
func (s *Service) CreateReport(c echo.Context) error {
	var reportSlice []Report

	err := c.Bind(&reportSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	// дата из хедера запроса
	createdAtHeader := c.Request().Header.Get("Date")
	var createdAt time.Time
	if createdAtHeader != "" {
		createdAt, err = time.Parse(time.RFC3339, createdAtHeader)
		if err != nil {
			s.logger.Error(err)
			return c.JSON(s.NewError(InvalidParams))
		}
	}

	repo := s.reportsRepo
	for _, report := range reportSlice {
		if createdAtHeader == "" {
			err = repo.CreateNewReport(report.Title, report.Description, time.Now())
		} else {
			err = repo.CreateNewReport(report.Title, report.Description, createdAt)
		}
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// UpdateReport обновляем репорт
func (s *Service) UpdateReport(c echo.Context) error {
	var reportSlice []Report
	err := c.Bind(&reportSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	// дата из хедера запроса
	createdAtHeader := c.Request().Header.Get("Date")
	var createdAt time.Time
	if createdAtHeader != "" {
		createdAt, err = time.Parse(time.RFC3339, createdAtHeader)
		if err != nil {
			s.logger.Error(err)
			return c.JSON(s.NewError(InvalidParams))
		}
	}

	repo := s.reportsRepo
	for _, report := range reportSlice {
		if createdAtHeader == "" {
			err = repo.UpdateReportById(id, report.Title, report.Description, time.Now())
		} else {
			err = repo.UpdateReportById(id, report.Title, report.Description, createdAt)
		}
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}

// DeleteReport удаляем из базы данных репорт
func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.DeleteReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}
