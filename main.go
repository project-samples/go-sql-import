package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	"github.com/core-go/log"
	"github.com/core-go/log/rotatelogs"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	err := config.Load(&cfg, "configs/config")
	if err != nil {
		panic(err)
	}
	log.Initialize(cfg.Log, rotatelogs.GetWriter)
	ctx := context.Background()
	log.Info(ctx, "Import file")
	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Errorf(ctx, "Error when initialize: %v", err)
		panic(err)
	}

	total, success, err := app.Import(ctx)
	fmt.Println(fmt.Sprintf("total: %d, success: %d", total, success))
	if err != nil {
		log.Errorf(ctx, "Error when import: %v", err)
		panic(err)
	}
	log.Info(ctx, "Imported file")
}
