package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Port:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

type Server struct {
	*http.Server
	config *Config
}

func NewServer(handler http.Handler, config *Config) *Server {
	if config == nil {
		config = DefaultConfig()
	}

	srv := &Server{
		Server: &http.Server{
			Addr:         config.Port,
			Handler:      handler,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
			IdleTimeout:  config.IdleTimeout,
		},
		config: config,
	}

	return srv
}

func (s *Server) Start() error {
	logger.Info("Starting server", zap.String("port", s.config.Port))
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	logger.Info("Shutting down server...")
	if err := s.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown:", zap.Error(err))
		return err
	}

	logger.Info("Server exited gracefully")
	return nil
}
