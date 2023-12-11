package web

import "github.com/google/uuid"

type InputSaveTransactionHistory struct {
	Id         uuid.UUID `json:"id"`
	From       string    `json:"from"`
	To         string    `json:"to" `
	Amount     int64     `validate:"required" json:"amount" `
	Fee        int64     `json:"fee" `
	Note       string    `json:"note" `
	Username   string    `validate:"required" json:"username" `
	MethodId   string    `validate:"required" json:"method_id" `
	AccountId  string    `validate:"required" json:"account_id" `
	CategoryId string    `validate:"required" json:"category_id" `
}

type InputEditTransactionHistory struct {
	Id         uuid.UUID `validate:"required" json:"id"`
	From       string    `json:"from"`
	To         string    `json:"to" `
	Amount     int64     `validate:"required" json:"amount" `
	Fee        int64     `json:"fee" `
	Note       string    `json:"note" `
	Username   string    `validate:"required" json:"username" `
	MethodId   string    `validate:"required" json:"method_id" `
	AccountId  string    `validate:"required" json:"account_id" `
	CategoryId string    `validate:"required" json:"category_id" `
}

type InputDeleteTransactionHistory struct {
	Id uuid.UUID `validate:"required" json:"id"`
}

type ReportByMethodIdInput struct {
	Username string `validate:"required" json:"username" `
	MethodId string `validate:"required" json:"method_id" `
}
