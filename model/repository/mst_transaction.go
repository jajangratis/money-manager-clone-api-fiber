package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type MstTransactionRepository interface {
	FindAll(tx *gorm.DB) []domain.MstTransaction
}
