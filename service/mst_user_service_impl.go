package service

import (
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api-fiber/app"
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/repository"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
	"gorm.io/gorm"
)

type MstUserServiceImpl struct {
	MstUserRepository repository.MstUserRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewMstUserService(mstUserRepository repository.MstUserRepository, DB *gorm.DB, validate *validator.Validate) MstUserService {
	return &MstUserServiceImpl{
		MstUserRepository: mstUserRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *MstUserServiceImpl) Login(login *web.LoginRequest) web.WebResponse {
	err := service.Validate.Struct(login)
	if err != nil {
		return helper.InvalidParameter()
	}
	tx := app.OpenConnection()
	data := service.MstUserRepository.FindOne(tx, login.Username)
	decodedDbPassword := helper.Decoded(data.Password)
	if login.Password != decodedDbPassword {
		return helper.InvalidParameterWithMessage("invalid_email_password")
	} else {
		token, err := helper.CreateJWT("userId", data.Username)
		if err != nil {
			helper.PanicIfError(err)
		}
		return helper.Ok(token)
	}
}
