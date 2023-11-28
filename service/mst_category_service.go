package service

import (
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
)

type MstCategoryService interface {
	FindAll() web.WebResponse
}
