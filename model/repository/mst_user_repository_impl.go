package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type MstUserRepositoryImpl struct{}

func NewMstUserRepositoryImpl() *MstUserRepositoryImpl {
	return &MstUserRepositoryImpl{}
}

func (repository *MstUserRepositoryImpl) FindOne(tx *gorm.DB, username string) domain.MstUser {
	defer func() {
		dbInstance, _ := tx.DB()
		_ = dbInstance.Close()
	}()
	var user domain.MstUser
	result := tx.First(&user, "username = ?", username).Error
	helper.PanicIfError(result)
	return user
}
