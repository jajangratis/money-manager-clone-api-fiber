package repository

import (
	"context"
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrTransactionHistoryRepositoryImpl struct{}

func NewTrTransactionHistoryRepositoryImpl() *TrTransactionHistoryRepositoryImpl {
	return &TrTransactionHistoryRepositoryImpl{}
}

func (repository *TrTransactionHistoryRepositoryImpl) FindAll(db *gorm.DB, username string) []domain.TrTransactionHistory {
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
	var data []domain.TrTransactionHistory
	result := db.Find(&data, "username = ?", username).Error
	helper.PanicIfError(result)
	return data
}

func (repository *TrTransactionHistoryRepositoryImpl) Save(ctx context.Context, db *gorm.DB, data domain.TrTransactionHistory) {
	err := db.Transaction(func(tx *gorm.DB) error {
		//err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = $1", "2").Error
		defer func() {
			dbInstance, _ := db.DB()
			_ = dbInstance.Close()
		}()
		result := db.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).WithContext(ctx).Save(&data).Error
		return result
	})
	if err != nil {
		helper.PanicIfError(err)
	}
}

func (repository *TrTransactionHistoryRepositoryImpl) SumData(ctx context.Context, db *gorm.DB, methodId string, username string) domain.AggregationResult {
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
	var result domain.AggregationResult
	err := db.Where("username = ? and method_id = ?", username, methodId).Model(&domain.TrTransactionHistory{}).Select("SUM(amount) as total_amount", "SUM(fee) as total_fee").Group("method_id").Find(&result).Error
	helper.PanicIfError(err)
	return result
}
