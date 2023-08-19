package jwt

import (
	"HOPE-backend/config"
	"HOPE-backend/pkg/jwt"
	"HOPE-backend/v3/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthorizeToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var res model.Response

		const bearerSchema = "Bearer "
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			res.Error = "authorization header not given"
			return c.JSON(http.StatusUnauthorized, res)
		}

		tokenString := authHeader[len(bearerSchema):]
		claim, err := jwt.AuthorizeToken(tokenString, config.Get().Server.SecretKey, false)
		if err != nil {
			res.Error = err.Error()
			return c.JSON(http.StatusUnauthorized, res)
		}

		c.Set("id", claim.Id)
		c.Set("role", claim.Role)
		c.Set("isVerified", claim.IsVerified)

		return next(c)
	}
}

func AuthorizeRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var res model.Response

			if c.Get("role").(string) != role {
				res.Error = "you do not have permission to access this"
				return c.JSON(http.StatusForbidden, res)
			}

			return next(c)
		}
	}
}
