package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
	"github.com/jajangratis/money-manager-clone-api-fiber/service"
)

type MstUserControllerImpl struct {
	MstUserService service.MstUserService
}

func NewMstUserController(mstUserService service.MstUserService) MstUserController {
	return &MstUserControllerImpl{
		MstUserService: mstUserService,
	}
}

func (controller *MstUserControllerImpl) Login(ctx *fiber.Ctx) error {
	request := new(web.LoginRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.JSON(helper.InvalidParameter())
	}
	data := controller.MstUserService.Login(request.Username, request.Password)
	return ctx.JSON(data)
}
