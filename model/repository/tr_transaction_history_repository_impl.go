package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type TrTransactionHistoryRepositoryImpl struct{}

func NewTrTransactionHistoryRepositoryImpl() *TrTransactionHistoryRepositoryImpl {
	return &TrTransactionHistoryRepositoryImpl{}
}

func (repository *TrTransactionHistoryRepositoryImpl) FindAll(tx *gorm.DB, username string) []domain.TrTransactionHistory {
	defer func() {
		dbInstance, _ := tx.DB()
		_ = dbInstance.Close()
	}()
	var data []domain.TrTransactionHistory
	result := tx.Find(&data, "username = ?", username).Error
	helper.PanicIfError(result)
	return data
}
