package service

import (
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
	"gorm.io/gorm"
)

type MstCategoryServiceImpl struct {
	MstCategoryRepository repository.MstCategoryRepository
	DB                    *gorm.DB
	Validate              *validator.Validate
}

func NewMstCategoryService(mstCategoryRepository repository.MstCategoryRepository, DB *gorm.DB, validate *validator.Validate) MstCategoryService {
	return &MstCategoryServiceImpl{
		MstCategoryRepository: mstCategoryRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *MstCategoryServiceImpl) FindAll() web.WebResponse {
	//var data []domain.MstCategory
	tx := app.OpenConnection()
	data := service.MstCategoryRepository.FindAll(tx)
	return web.WebResponse{
		Code:    200,
		Status:  "success",
		Message: "ok",
		Data:    data,
	}
}
