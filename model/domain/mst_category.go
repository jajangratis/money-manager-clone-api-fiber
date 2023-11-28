package domain

import (
	"github.com/google/uuid"
	"time"
)

// ALIAS Category
type MstCategory struct {
	Id                 uuid.UUID              `json:"id" gorm:"column:id"`
	CategoryId         string                 `json:"category_id" gorm:"primary_key;column:category_id"`
	CategoryName       string                 `json:"category_name" gorm:"column:category_name"`
	MethodId           string                 `json:"method_id" gorm:"column:method_id"`
	CreatedDate        time.Time              `json:"created_date" gorm:"column:created_date;autoCreateTime"`
	UpdatedDate        time.Time              `json:"updated_date" gorm:"column:updated_date;autoCreateTime"`
	TransactionHistory []TrTransactionHistory `json:"transaction_history" gorm:"foreignKey:category_id;references:category_id"`
	Transaction        *MstTransaction        `json:"transaction" gorm:"foreignKey:method_id;references:method_id"`
}

func (u *MstCategory) TableName() string {
	return "mst_category"
}
