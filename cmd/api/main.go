package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/yvv4git/userinfo/internal/domain/usecases"
	"github.com/yvv4git/userinfo/internal/infrastructure/api"
	"github.com/yvv4git/userinfo/internal/infrastructure/config"
	"github.com/yvv4git/userinfo/internal/infrastructure/logger"
	"github.com/yvv4git/userinfo/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	defaultLogger := logger.SetupDefaultLogger()

	cfg := config.NewConfig(defaultLogger)
	if err := cfg.Load(); err != nil {
		defaultLogger.Error("load config", "err", err)
	}

	log := logger.SetupLoggerWithLevel(cfg.LogLevel)
	log.Info("Service web-api started")
	defer log.Info("Service web-api shutdown")

	ucConsoleProcessor := usecases.NewDstConsole(log)
	svcUserInfo := service.NewUserInfo(ucConsoleProcessor)
	webSrv := api.NewWeb(log, &cfg.API, svcUserInfo)
	if err := webSrv.Run(ctx); err != nil {
		log.Error("failed to run web server", "err", err)
	}
}
