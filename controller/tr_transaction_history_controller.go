package controller

import "github.com/gofiber/fiber/v2"

type TrTransactionHistoryController interface {
	FindAll(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	ReportIncomeExpense(ctx *fiber.Ctx) error
}
