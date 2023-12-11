package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (controller *TrTransactionHistoryControllerImpl) Edit(ctx *fiber.Ctx) error {
	username := ctx.Get("userId")
	data := new(web.InputEditTransactionHistory)
	err := ctx.BodyParser(data)
	if err != nil {
		return err
	}
	data.Username = username
	result := controller.TrTransactionHistoryService.Edit(data)
	return ctx.JSON(result)
}

func (controller *TrTransactionHistoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := new(web.InputDeleteTransactionHistory)
	err := ctx.BodyParser(data)
	if err != nil {
		return err
	}
	// Parse the string into a UUID
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return err
	}
	data.Id = parsedUUID
	result := controller.TrTransactionHistoryService.Delete(data)
	return ctx.JSON(result)
}
