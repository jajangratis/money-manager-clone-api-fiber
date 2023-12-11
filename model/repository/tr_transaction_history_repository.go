package repository

import (
	"context"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type TrTransactionHistoryRepository interface {
	FindAll(tx *gorm.DB, username string) []domain.TrTransactionHistory
	Save(ctx context.Context, db *gorm.DB, data domain.TrTransactionHistory)
	SumData(ctx context.Context, db *gorm.DB, methodId string, username string) domain.AggregationResult
	Edit(ctx context.Context, db *gorm.DB, data domain.TrTransactionHistory)
	Delete(ctx context.Context, db *gorm.DB, data domain.TrTransactionHistory)
}
