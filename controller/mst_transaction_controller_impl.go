package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/service"
)

type MstTransactionControllerImpl struct {
	MstTransactionService service.MstTransactionService
}

func NewMstTransactionControllerImpl(mstTransactionService service.MstTransactionService) MstTransactionController {
	return &MstTransactionControllerImpl{
		MstTransactionService: mstTransactionService,
	}
}

func (controller *MstTransactionControllerImpl) FindAll(ctx *fiber.Ctx) error {
	data := controller.MstTransactionService.FindAll()
	return ctx.JSON(data)
}
