package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test_1/internal/repository"
	"test_1/internal/server"
	service2 "test_1/internal/service"
)

func main() {
	logger := logrus.New()

	cfg, err := LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("failed to load config")
	}

	repo := repository.New()

	service := service2.NewService(repo)

	ginRouter := createGINRouter()

	err = server.RegisterControllers(ginRouter.Group("/v1"), service)

	if err != nil {
		logger.WithError(err).Fatal("failed to register controller")
	}

	server := http.Server{
		Addr:    cfg.Address + ":" + cfg.Port,
		Handler: ginRouter,
	}

	go func() {
		logger.Info("start book service")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("debug listen and serve error: %+v", err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	logger.Info("stop book service")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logger.WithError(err).Fatal("failed to shutdown http server")
	}
}
