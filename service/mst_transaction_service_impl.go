package service

import (
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
	"gorm.io/gorm"
)

type MstTransactionServiceImpl struct {
	MstTransactionRepository repository.MstTransactionRepository
	DB                       *gorm.DB
	Validate                 *validator.Validate
}

func NewMstTransactionServiceImpl(mstTransactionRepository repository.MstTransactionRepository, DB *gorm.DB, validate *validator.Validate) MstTransactionService {
	return &MstTransactionServiceImpl{
		MstTransactionRepository: mstTransactionRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *MstTransactionServiceImpl) FindAll() web.WebResponse {
	//var data []domain.MstCategory
	tx := app.OpenConnection()
	data := service.MstTransactionRepository.FindAll(tx)
	return web.WebResponse{
		Code:    200,
		Status:  "success",
		Message: "ok",
		Data:    data,
	}
}
