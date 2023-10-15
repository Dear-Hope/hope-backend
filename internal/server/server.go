package server

import (
	"HOPE-backend/internal/api/auth"
	"HOPE-backend/internal/api/consultation"
	"HOPE-backend/internal/api/expert"
	"HOPE-backend/internal/api/health"
	"HOPE-backend/internal/api/user"
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/middleware/jwt"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	router        *echo.Echo
	HealthHandler *health.Handler
	AuthHandler   *auth.Handler
	UserHandler   *user.Handler
	ExpertHandler *expert.Handler
	ConsulHandler *consultation.Handler
}

func (s *Server) serve(port string) error {
	s.router = echo.New()
	s.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: middleware.DefaultCORSConfig.AllowOrigins,
		AllowHeaders: []string{"*"},
		AllowMethods: append(middleware.DefaultCORSConfig.AllowMethods, http.MethodOptions),
	}))

	s.router.Use(middleware.Logger())

	// handle static
	s.router.Static("/assets", "./assets")

	// Register health handler
	s.router.GET("/server/health", s.HealthHandler.Check, jwt.AuthorizeToken, jwt.AuthorizeRole("NORMAL"))

	// Register api group - backward compatibility with previous version
	apiV3 := s.router.Group("/api/v3")

	// TODO: Register all endpoint below

	// Register auth handler
	authRouter := apiV3.Group("/auth")
	authRouter.POST("/login", s.AuthHandler.Login)
	authRouter.POST("/login/refresh", s.AuthHandler.RefreshToken)
	authRouter.POST("/resend", s.AuthHandler.ResendOtp)
	authRouter.POST("/password/reset", s.AuthHandler.ResetPassword)
	authRouter.POST("/password/change", s.AuthHandler.ChangePassword)

	// Register user handler
	userRouter := apiV3.Group("/user")
	userRouter.POST("/register", s.UserHandler.Register)
	userRouter.POST("/activate", s.UserHandler.Verify)
	userRouter.GET("/me", s.UserHandler.GetUserMe, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.UserRole))
	userRouter.PUT("/me", s.UserHandler.UpdateUserMe, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.UserRole))
	userRouter.POST("/me/upload/photo", s.UserHandler.UploadProfilePhoto, jwt.AuthorizeToken,
		jwt.AuthorizeRole(constant.UserRole))

	// Register expert handler
	expertRouter := apiV3.Group("/expert")
	expertRouter.POST("/register", s.ExpertHandler.Register)
	expertRouter.GET("/me", s.ExpertHandler.GetExpertMe, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.ExpertRole))
	expertRouter.PUT("/me", s.ExpertHandler.UpdateExpertMe, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.ExpertRole))
	expertRouter.GET("/schedule", s.ExpertHandler.GetSchedule, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.ExpertRole))
	expertRouter.PUT("/schedule", s.ExpertHandler.UpdateSchedule, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.ExpertRole))
	expertRouter.GET("/:id/schedule", s.ExpertHandler.GetScheduleUser, jwt.AuthorizeToken,
		jwt.AuthorizeRole(constant.UserRole))
	expertRouter.GET("/consultation", s.ExpertHandler.ListConsultation, jwt.AuthorizeToken,
		jwt.AuthorizeRole(constant.ExpertRole))
	expertRouter.GET("/consultation/:id", s.ExpertHandler.DetailConsultation, jwt.AuthorizeToken,
		jwt.AuthorizeRole(constant.ExpertRole))
	expertRouter.GET("/review", s.ExpertHandler.GetReview, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.ExpertRole))

	// Register consultation handler
	consulRouter := apiV3.Group("/consultation")
	consulRouter.PUT("/:id/status", s.ConsulHandler.UpdateStatus, jwt.AuthorizeToken, jwt.AuthorizeRole(constant.ExpertRole))

	return s.router.Start(":" + port)
}

func (s *Server) gracefulStop(timeout int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := s.router.Shutdown(ctx); err != nil {
		log.Fatalf("error graceful shutdown: %v", err)
	}
}

func (s *Server) Run(port string, timeout int64) error {
	idleConnClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// Notify when received terminate signal
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals

		// We received an os signal, shut down.
		s.gracefulStop(timeout)
		log.Println("Server shutdown gracefully")
		close(idleConnClosed)
	}()

	log.Println("Server running on port", port)
	if err := s.serve(port); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return err
	}

	<-idleConnClosed
	log.Println("Server stopping")
	return nil
}
