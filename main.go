package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"login-wrapper/app"
	"login-wrapper/config"
	"login-wrapper/pkg/logging"
	"login-wrapper/pkg/server"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	log := logging.NewLogger(os.Getenv("LOG_LEVEL"), os.Getenv("ENVIRONMENT"))
	ctx = logging.WithLogger(ctx, log)

	defer func() {
		done()
		if r := recover(); r != nil {
			log.Errorf("application went wrong. Panic err=%v", r)
		}
	}()

	err := realMain(ctx)
	done()
	if err != nil {
		log.Errorf("realMain has failed with err=%v", err)
		return
	}

	log.Info("APP shutdown successful")
}

func realMain(ctx context.Context) error {
	log := logging.FromContext(ctx)

	cfg, err := config.LoadFromEnv(ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("load config from environment failed with err=%v", err))
	}
	srv, err := server.New(cfg.HTTPPort)
	if err != nil {
		return err
	}

	log.Infof("HTTP Server running on PORT: %s", cfg.HTTPPort)

	return srv.ServeHTTPHandler(ctx, app.New().Routes(ctx))
}
