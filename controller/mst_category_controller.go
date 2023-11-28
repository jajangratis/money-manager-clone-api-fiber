package controller

import "github.com/gofiber/fiber/v2"

type MstCategoryController interface {
	FindAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}
