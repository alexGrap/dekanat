package usecase

import (
	"github.com/gofrs/uuid"
	"ozon/internal/models"
	"time"
)

func GetMarks(body models.RestBody) (models.RestResult, error) {
	switch body.LastName {
	case "Карамышев":
		return models.RestResult{LastName: body.LastName, Stats: "Лучший студент"}, nil
	case "Паргасов":
		return models.RestResult{LastName: body.LastName, Stats: "Неплохой студент"}, nil
	case "Сомов":
		return models.RestResult{LastName: body.LastName, Stats: "Хитрый жук"}, nil
	default:
		return models.RestResult{LastName: body.LastName, Stats: "Хороший малый"}, nil

	}
}

func GetComing(body models.RestBody) (models.RestResult, error) {
	switch body.LastName {
	case "Карамышев":
		return models.RestResult{LastName: body.LastName, Stats: "10/10"}, nil
	case "Паргасов":
		return models.RestResult{LastName: body.LastName, Stats: "100/10"}, nil
	case "Сомов":
		return models.RestResult{LastName: body.LastName, Stats: "5/10"}, nil
	default:
		return models.RestResult{LastName: body.LastName, Stats: "8/10"}, nil

	}
}

func GetSchedule(body models.ScheduleRequest) (models.ScheduleResponse, error) {
	switch body.Date {
	case "20/05/2024":
		return models.ScheduleResponse{Lessons: []*models.Lesson{{ID: uuid.UUID{}, Date: time.Date(2024, 5, 20, 12, 20, 0, 0, time.Local), Name: "Информационные системы и технологии", Type: "Лабораторная работа", Professor: "Бычков С.Ю.", Class: "450"}}}, nil

	default:
		return models.ScheduleResponse{}, nil
	}
}
