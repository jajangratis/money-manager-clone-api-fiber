package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type MstTransactionRepositoryImpl struct{}

func NewMstTransactionRepositoryImpl() *MstTransactionRepositoryImpl {
	return &MstTransactionRepositoryImpl{}
}

func (repository *MstTransactionRepositoryImpl) FindAll(tx *gorm.DB) []domain.MstTransaction {
	defer func() {
		dbInstance, _ := tx.DB()
		_ = dbInstance.Close()
	}()
	var mstTransaction []domain.MstTransaction

	result := tx.Preload("Category").Find(&mstTransaction).Error
	helper.PanicIfError(result)
	return mstTransaction
}
