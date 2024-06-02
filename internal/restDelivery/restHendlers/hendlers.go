package restHendlers

import (
	"github.com/gofiber/fiber/v2"
	"ozon/internal/models"
	"ozon/internal/usecase"
)

func GetMarks(ctx *fiber.Ctx) error {
	var (
		body models.RestBody
	)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.JSON(err)
	}
	result, err := usecase.GetMarks(body)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(result)
}

func GetComing(ctx *fiber.Ctx) error {
	var (
		body models.RestBody
	)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.JSON(err)
	}
	result, err := usecase.GetComing(body)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(result)
}

func GetSchedule(ctx *fiber.Ctx) error {
	var body models.ScheduleRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.JSON(err)
	}
	response, err := usecase.GetSchedule(body)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(response)
}
