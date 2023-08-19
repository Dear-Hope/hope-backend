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
	Login(ctx context.Context, req auth.LoginRequest) (*auth.TokenPairResponse, *response.ServiceError)
	ResendOtp(ctx context.Context, email string) *response.ServiceError
	ResetPassword(ctx context.Context, email string) *response.ServiceError
	ChangePassword(ctx context.Context, req auth.ChangePasswordRequest) (*auth.TokenPairResponse, *response.ServiceError)
	RefreshToken(ctx context.Context, token string) (*auth.TokenPairResponse, *response.ServiceError)
}

type Handler struct {
	svc service
}

func New(svc service) *Handler {
	return &Handler{svc: svc}
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

func (h *Handler) ResendOtp(c echo.Context) error {
	var (
		res response.Response
		req = map[string]string{}
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	svcErr := h.svc.ResendOtp(c.Request().Context(), req["email"])
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.ResendOtp]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"success": true}
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) ResetPassword(c echo.Context) error {
	var (
		res response.Response
		req = map[string]string{}
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	svcErr := h.svc.ResetPassword(c.Request().Context(), req["email"])
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.ResetPassword]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	res.Result = map[string]bool{"success": true}
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) ChangePassword(c echo.Context) error {
	var (
		res    response.Response
		req    auth.ChangePasswordRequest
		svcErr *response.ServiceError
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.ChangePassword(c.Request().Context(), req)
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.ChangePassword]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var (
		res    response.Response
		req    = map[string]string{}
		svcErr *response.ServiceError
	)

	err := c.Bind(&req)
	if err != nil {
		res.Error = fmt.Sprintf("invalid parameters: %s", err.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Result, svcErr = h.svc.RefreshToken(c.Request().Context(), req["refresh"])
	if svcErr != nil {
		c.Logger().Errorf("[AuthHandler.RefreshToken]%v", svcErr.Err)
		res.Error = svcErr.Msg
		return c.JSON(svcErr.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}
