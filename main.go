package main

import (
	"context"
	"fmt"

	"github.com/core-go/config"
	"github.com/core-go/log/rotatelogs"
	"github.com/core-go/log/zap"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	err := config.Load(&cfg, "configs/config")
	if err != nil {
		panic(err)
	}
	log.InitializeWithWriter(cfg.Log, rotatelogs.GetWriter)
	ctx := context.Background()
	log.Info(ctx, "Import file")
	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Errorf(ctx, "Error when initialize: %v", err)
		panic(err)
	}

	total, success, err := app.Import(ctx)
	if err != nil {
		log.Errorf(ctx, "Error when import: %v", err)
		panic(err)
	}
	log.Info(ctx, fmt.Sprintf("total: %d, success: %d", total, success))
	log.Info(ctx, "Imported file")
}
