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
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
	log.Initialize(conf.Log, rotatelogs.GetWriter)
	fmt.Println("Import file")
	ctx := context.Background()
	log.Info(ctx, "Import file")
	app, err := app.NewApp(ctx, conf)
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
