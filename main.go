package main

import (
	"context"
	"fmt"
	"github.com/bpsaunders/user-api/config"
	"github.com/bpsaunders/user-api/handlers"
	"github.com/bpsaunders/user-api/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Error(fmt.Sprintf("error configuring service: %s. Exiting", err))
		os.Exit(1)
	}

	setLogLevel(cfg)

	userService := service.NewUserService(cfg)
	mainRouter := mux.NewRouter()

	handlers.Register(mainRouter, userService)

	h := &http.Server{
		Addr:    ":8888",
		Handler: mainRouter,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// run server in new go routine to allow app shutdown signal wait below
	go func() {
		err := h.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			os.Exit(1)
		}
	}()

	// wait for app shutdown message before attempting to close server gracefully
	<-stop

	log.Info("shutting down server...")

	userService.Shutdown()
	timeout := time.Duration(5) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = h.Shutdown(ctx)
	if err != nil {
		log.Error(fmt.Errorf("failed to shutdown gracefully: %s", err))
	} else {
		log.Info("server shutdown gracefully")
	}
}

func setLogLevel(cfg *config.Config) {

	if cfg.LogLevel != "" {
		log.Info(fmt.Sprintf("Log level set in environment, attempting to set log level to: %s", cfg.LogLevel))
		lvl, err := log.ParseLevel(cfg.LogLevel)
		if err != nil {
			log.Error(fmt.Sprintf("failed to set log level: %s. Exiting", err))
			os.Exit(1)
		}
		log.SetLevel(lvl)
		log.Info("Log level set successfully")
	}
}
