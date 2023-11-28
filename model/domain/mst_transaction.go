package domain

import (
	"github.com/google/uuid"
	"time"
)

//ALIAS TransactionMethod

type MstTransaction struct {
	Id                 uuid.UUID              `json:"id" gorm:"column:id"`
	MethodId           string                 `json:"method_id" gorm:"primary_key;column:method_id"`
	MethodName         string                 `json:"method_name" gorm:"column:method_name"`
	CreatedDate        time.Time              `json:"created_date" gorm:"column:created_date;autoCreateTime"`
	UpdatedDate        time.Time              `json:"updated_date" gorm:"column:updated_date;autoCreateTime"`
	TransactionHistory []TrTransactionHistory `json:"transaction_history" gorm:"foreignKey:method_id;references:method_id"`
	Category           []MstCategory          `json:"category" gorm:"foreignKey:method_id;references:method_id"`
}

func (u *MstTransaction) TableName() string {
	return "mst_transaction_method"
}
