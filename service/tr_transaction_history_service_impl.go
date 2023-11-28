package service

import (
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
	"gorm.io/gorm"
)

type TrTransactionServiceImpl struct {
	TrTransactionHistoryRepository repository.TrTransactionHistoryRepository
	DB                             *gorm.DB
	Validate                       *validator.Validate
}

func NewTrTransactionServiceImpl(trTransactionHistoryRepository repository.TrTransactionHistoryRepository, DB *gorm.DB, validate *validator.Validate) TrTransactionHistoryService {
	return &TrTransactionServiceImpl{
		TrTransactionHistoryRepository: trTransactionHistoryRepository,
		DB:                             DB,
		Validate:                       validate,
	}
}

func (service *TrTransactionServiceImpl) FindAll(username string) web.WebResponse {
	if username == "" {
		return helper.InvalidParameter()
	}
	tx := app.OpenConnection()
	data := service.TrTransactionHistoryRepository.FindAll(tx, username)
	return helper.Ok(data)
}
