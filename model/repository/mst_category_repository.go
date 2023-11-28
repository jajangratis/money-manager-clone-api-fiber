package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type MstCategoryRepository interface {
	FindAll(tx *gorm.DB) []domain.MstCategory
}
