package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Missing or invalid token")
		}

		claims, err := ParseToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
		}

		c.Set("userID", claims.ID)
		return next(c)
	}
}
