package domain

import (
	"github.com/google/uuid"
	"time"
)

// ALIAS User

type MstUser struct {
	Id                 uuid.UUID              `json:"id" gorm:"column:id"`
	Username           string                 `json:"username" gorm:"primary_key;column:username"`
	Email              string                 `json:"email" gorm:"column:email"`
	Password           string                 `json:"password" gorm:"column:password"`
	CreatedDate        time.Time              `json:"created_date" gorm:"column:created_date;autoCreateTime"`
	UpdatedDate        time.Time              `json:"updated_date" gorm:"column:updated_date;autoCreateTime"`
	TransactionHistory []TrTransactionHistory `json:"transaction_history" gorm:"foreignKey:username;references:username"`
}

func (u *MstUser) TableName() string {
	return "mst_user"
}
