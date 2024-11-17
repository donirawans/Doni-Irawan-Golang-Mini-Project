package utils

import (
	"github.com/labstack/echo/v4"
)

func JSONErrorResponse(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode, map[string]interface{}{
		"success": false,
		"error":   message,
	})
}
