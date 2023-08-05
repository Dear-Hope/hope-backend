package server

import (
	"HOPE-backend/internal/api/auth"
	"HOPE-backend/internal/api/health"
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

	// Register api group
	api := s.router.Group("/api")

	// TODO: Register all endpoint below

	// Register auth handler
	authRouter := api.Group("/auth")
	authRouter.POST("/register", s.AuthHandler.Register)

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
