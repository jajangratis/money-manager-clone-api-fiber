package service

import "github.com/jajangratis/money-manager-clone-api-fiber/model/web"

type TrTransactionHistoryService interface {
	FindAll(username string) web.WebResponse
}
