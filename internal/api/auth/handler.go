package auth

import (
	"HOPE-backend/internal/entity/auth"
	"HOPE-backend/internal/entity/response"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type service interface {
	Register(ctx context.Context, req auth.RegisterRequest) (*auth.TokenPairResponse, *response.ServiceError)
	Login(ctx context.Context, req auth.LoginRequest) (*auth.TokenPairResponse, *response.ServiceError)
}

type Handler struct {
	svc service
}

func New(svc service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(c echo.Context) error {
	var (
		res    response.Response
		req    auth.RegisterRequest
		svcErr *response.ServiceError
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.Register(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.Register]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c echo.Context) error {
	var (
		req    auth.LoginRequest
		res    response.Response
		svcErr *response.ServiceError
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.Login(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.Login]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
