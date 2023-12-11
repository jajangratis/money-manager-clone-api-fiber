package service

import "github.com/jajangratis/money-manager-clone-api-fiber/model/web"

type TrTransactionHistoryService interface {
	FindAll(username string) web.WebResponse
	Save(data *web.InputSaveTransactionHistory) web.WebResponse
	ReportByMethodId(input *web.ReportByMethodIdInput) web.WebResponse
	Edit(data *web.InputEditTransactionHistory) web.WebResponse
	Delete(data *web.InputDeleteTransactionHistory) web.WebResponse
}
