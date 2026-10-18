package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-gin-starter/internal/config"
	"go-gin-starter/internal/database"
	"go-gin-starter/internal/logger"
	"go-gin-starter/internal/redis"
	"go-gin-starter/internal/router"

	"github.com/gin-gonic/gin"
)

var (
	configPath string
	Version    = "dev"
	BuildTime  = "unknown"
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yaml", "config file path")
}

func main() {
	flag.Parse()

	if err := config.Init(configPath); err != nil {
		fmt.Printf("failed to init config: %v\n", err)
		os.Exit(1)
	}

	cfg := config.GetConfig()

	if err := logger.Init(); err != nil {
		fmt.Printf("failed to init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	if err := database.Init(); err != nil {
		logger.Fatalf("failed to init database: %v", err)
	}
	defer database.Close()

	if err := redis.Init(); err != nil {
		logger.Fatalf("failed to init redis: %v", err)
	}
	defer redis.Close()

	gin.SetMode(cfg.Server.Mode)

	r := gin.New()
	router.Setup(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	logger.Infof("starting %s version=%s build=%s on port %d", cfg.Server.AppName, Version, BuildTime, cfg.Server.Port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server forced to shutdown: %v", err)
	}

	logger.Info("server exited")
}