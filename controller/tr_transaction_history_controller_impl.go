package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/service"
)

type TrTransactionHistoryControllerImpl struct {
	TrTransactionHistoryService service.TrTransactionHistoryService
}

func NewTrTransactionHistoryControllerImpl(trTransactionHistoryService service.TrTransactionHistoryService) TrTransactionHistoryController {
	return &TrTransactionHistoryControllerImpl{
		TrTransactionHistoryService: trTransactionHistoryService,
	}
}

func (controller *TrTransactionHistoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	username := ctx.Get("userId")
	data := controller.TrTransactionHistoryService.FindAll(username)
	return ctx.JSON(data)
}
