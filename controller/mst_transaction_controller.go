package controller

import "github.com/gofiber/fiber/v2"

type MstTransactionController interface {
	FindAll(ctx *fiber.Ctx) error
}
