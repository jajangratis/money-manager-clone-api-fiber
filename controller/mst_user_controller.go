package controller

import "github.com/gofiber/fiber/v2"

type MstUserController interface {
	Login(ctx *fiber.Ctx) error
}
