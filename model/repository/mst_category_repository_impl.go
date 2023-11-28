package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type MstCategoryRepositoryImpl struct{}

func NewMstCategoryRepositoryImpl() *MstCategoryRepositoryImpl {
	return &MstCategoryRepositoryImpl{}
}

func (repository *MstCategoryRepositoryImpl) FindAll(tx *gorm.DB) []domain.MstCategory {
	defer func() {
		dbInstance, _ := tx.DB()
		_ = dbInstance.Close()
	}()
	var mstCategory []domain.MstCategory
	result := tx.Find(&mstCategory).Error
	helper.PanicIfError(result)
	return mstCategory
}
