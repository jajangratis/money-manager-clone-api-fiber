package web

type InputSaveTransactionHistory struct {
	From       string `json:"from"`
	To         string `json:"to" `
	Amount     int64  `validate:"required" json:"amount" `
	Fee        int64  `json:"fee" `
	Note       string `json:"note" `
	Username   string `validate:"required" json:"username" `
	MethodId   string `validate:"required" json:"method_id" `
	AccountId  string `validate:"required" json:"account_id" `
	CategoryId string `validate:"required" json:"category_id" `
}

type ReportByMethodIdInput struct {
	Username string `validate:"required" json:"username" `
	MethodId string `validate:"required" json:"method_id" `
}
