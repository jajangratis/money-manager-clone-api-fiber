package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jajangratis/money-manager-clone-api-fiber/helper"
)

func convertToString(value interface{}) (string, error) {
	// Check if the underlying type is string
	stringValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("conversion to string failed: underlying type is not string")
	}

	return stringValue, nil
}

func TokenExtractor(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.JSON(helper.InvalidParameterWithMessage("no_token_provided"))
	} else {
		data, err := helper.ExtractJWTToken(token)
		if err != nil {
			return ctx.JSON(helper.InvalidParameterWithMessage("invalid_token"))
		} else {
			userId, _ := convertToString(data["userId"])
			ctx.Request().Header.Add("userId", userId)
			return ctx.Next()
		}
	}
}
