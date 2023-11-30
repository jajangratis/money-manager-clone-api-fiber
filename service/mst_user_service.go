package service

import "github.com/jajangratis/money-manager-clone-api-fiber/model/web"

type MstUserService interface {
	Login(login *web.LoginRequest) web.WebResponse
}
