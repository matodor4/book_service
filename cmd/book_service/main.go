package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	_ "test_1/docs"
	"test_1/internal/repository"
	"test_1/internal/server"
	"test_1/internal/service"
)

func main() {
	logger := logrus.New()

	cfg, err := LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("failed to load config")
	}

	repo := repository.New()

	bookService := service.NewService(repo)

	ginRouter := createGINRouter()

	err = server.RegisterControllers(ginRouter, bookService)

	if err != nil {
		logger.WithError(err).Fatal("failed to register controller")
	}

	addr := cfg.Address + ":" + cfg.Port

	server := http.Server{
		Addr:    addr,
		Handler: ginRouter,
	}

	go func() {
		logger.Info("start book bookService")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("debug listen and serve error: %+v", err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	logger.Info("stop book bookService")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logger.WithError(err).Fatal("failed to shutdown http server")
	}
}
