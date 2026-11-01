package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yjhss/newspaper-ops-monitor/internal/adapter/inbound/scheduler"
	fsadapter "github.com/yjhss/newspaper-ops-monitor/internal/adapter/outbound/filesystem"
	logadapter "github.com/yjhss/newspaper-ops-monitor/internal/adapter/outbound/logging"
	"github.com/yjhss/newspaper-ops-monitor/internal/application/service"
	"github.com/yjhss/newspaper-ops-monitor/internal/infrastructure/config"
)

func main() {

	logger := logadapter.NewLogger()

	appConfigPath := os.Getenv("APP_CONFIG")

	if appConfigPath == "" {
		appConfigPath = "./configs/app.yaml"
	}

	appConfig, err := config.LoadApp(appConfigPath)
	if err != nil {
		logger.Error(
			"failed to load application config",
			"error", err,
		)
		os.Exit(1)
	}

	profiles, err := config.LoadMediaDirectory(
		appConfig.MediaConfigDir,
	)

	if err != nil {
		logger.Error(
			"failed to load media configs",
			"error", err,
		)
		os.Exit(1)
	}

	logger.Info(
		"media configs loaded",
		"count", len(profiles),
	)

	fileScanner := fsadapter.NewScanner()

	monitorService := service.NewMonitorService(
		fileScanner,
		logger,
		profiles,
	)

	pollingScheduler := scheduler.NewPollingScheduler(
		monitorService,
		time.Duration(appConfig.PollIntervalSeconds)*time.Second,
	)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer cancel()

	logger.Info(
		"monitor started",
		"poll_interval_seconds",
		appConfig.PollIntervalSeconds,
	)

	pollingScheduler.Run(ctx)

	logger.Info("monitor stopped")
}
