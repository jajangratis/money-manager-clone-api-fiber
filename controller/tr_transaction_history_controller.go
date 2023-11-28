package controller

import "github.com/gofiber/fiber/v2"

type TrTransactionHistoryController interface {
	FindAll(ctx *fiber.Ctx) error
}
