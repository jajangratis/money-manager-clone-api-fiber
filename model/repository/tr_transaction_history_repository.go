package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type TrTransactionHistoryRepository interface {
	FindAll(tx *gorm.DB, username string) []domain.TrTransactionHistory
}
