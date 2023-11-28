package domain

import (
	"github.com/google/uuid"
	"time"
)

//ALIAS TransactionHistory

type TrTransactionHistory struct {
	Id                uuid.UUID       `json:"id" gorm:"primary_key;column:id"`
	TransactionMethod *MstTransaction `json:"transaction_method" gorm:"foreignKey:method_id;references:method_id"`
	AccountType       *MstAccountType `json:"account_type" gorm:"foreignKey:account_id;references:account_id"`
	Category          *MstCategory    `json:"category" gorm:"foreignKey:category_id;references:category_id"`
	User              *MstUser        `json:"user" gorm:"foreignKey:username;references:username"`
	From              string          `json:"from" gorm:"column:from"`
	To                string          `json:"to" gorm:"column:to"`
	Amount            int64           `json:"amount" gorm:"column:amount"`
	Fee               int64           `json:"fee" gorm:"column:fee"`
	Note              string          `json:"note" gorm:"column:note"`
	Username          string          `json:"username" gorm:"column:username"`
	MethodId          string          `json:"method_id" gorm:"column:method_id"`
	AccountId         string          `json:"account_id" gorm:"column:account_id"`
	CategoryId        string          `json:"category_id" gorm:"column:category_id"`
	CreatedDate       time.Time       `json:"created_date" gorm:"column:created_date;autoCreateTime"`
	UpdatedDate       time.Time       `json:"updated_date" gorm:"column:updated_date;autoCreateTime"`
}

func (u *TrTransactionHistory) TableName() string {
	return "tr_transaction_history"
}
