package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/model/web"
)

func InvalidParameter() web.WebResponse {
	return web.WebResponse{
		Code:    fiber.StatusBadRequest,
		Status:  "false",
		Message: "invalid_parameter",
	}
}

func Ok(data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "true",
		Message: "ok",
		Data:    data,
	}
}

func InvalidParameterWithMessage(msg string) web.WebResponse {
	return web.WebResponse{
		Code:    fiber.StatusBadRequest,
		Status:  "false",
		Message: msg,
	}
}
