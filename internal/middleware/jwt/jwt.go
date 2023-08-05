package jwt

import (
	"HOPE-backend/config"
	"HOPE-backend/v3/model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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
		token, err := ValidateToken(tokenString)
		if err != nil {
			res.Error = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if int64(claims["expires"].(float64)) < time.Now().Unix() {
				res.Error = "access token has expired"
				return c.JSON(http.StatusUnauthorized, res)
			}
			if !claims["isVerified"].(bool) {
				res.Error = "account has not been verified yet"
				return c.JSON(http.StatusUnauthorized, res)
			}
			userId := uint64(claims["userId"].(float64))
			role := claims["role"].(string)
			c.Set("userId", userId)
			c.Set("role", role)
			return next(c)
		} else {
			res.Error = "invalid token"
			return c.JSON(http.StatusUnauthorized, res)
		}
	}
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(config.Get().Server.SecretKey), nil
	})
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
