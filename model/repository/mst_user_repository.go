package repository

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/model/domain"
	"gorm.io/gorm"
)

type MstUserRepository interface {
	FindOne(tx *gorm.DB, username string) domain.MstUser
}
