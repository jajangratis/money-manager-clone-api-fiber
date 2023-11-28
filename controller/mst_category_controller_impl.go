package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/service"
)

type MstCategoryControllerImpl struct {
	MstCategoryService service.MstCategoryService
}

func NewMstCategoryController(mstCategoryService service.MstCategoryService) MstCategoryController {
	return &MstCategoryControllerImpl{
		MstCategoryService: mstCategoryService,
	}
}

func (controller MstCategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	data := controller.MstCategoryService.FindAll()
	return ctx.JSON(data)
}

func (controller MstCategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (controller MstCategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
