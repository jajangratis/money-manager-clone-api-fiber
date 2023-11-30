package domain

import (
	"github.com/google/uuid"
	"time"
)

//ALIAS AccountType

type MstAccountType struct {
	Id                 uuid.UUID            `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	AccountId          string               `json:"account_id" gorm:"primary_key;column:account_id"`
	AccountName        string               `json:"account_name" gorm:"column:account_name"`
	CreatedDate        time.Time            `json:"created_date" gorm:"column:created_date;autoCreateTime"`
	UpdatedDate        time.Time            `json:"updated_date" gorm:"column:updated_date;autoCreateTime"`
	TransactionHistory TrTransactionHistory `json:"transaction_history" gorm:"foreignKey:account_id;references:account_id"`
}

func (MstAccountType) TableName() string {
	return "mst_account_type"
}
