package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
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

func (controller *TrTransactionHistoryControllerImpl) Save(ctx *fiber.Ctx) error {
	username := ctx.Get("userId")
	data := new(web.InputSaveTransactionHistory)
	err := ctx.BodyParser(data)
	if err != nil {
		return err
	}
	data.Username = username
	result := controller.TrTransactionHistoryService.Save(data)
	return ctx.JSON(result)
}

func (controller *TrTransactionHistoryControllerImpl) ReportIncomeExpense(ctx *fiber.Ctx) error {
	username := ctx.Get("userId")
	data := new(web.ReportByMethodIdInput)
	err := ctx.BodyParser(data)
	if err != nil {
		return err
	}
	data.Username = username
	result := controller.TrTransactionHistoryService.ReportByMethodId(data)
	return ctx.JSON(result)
}
