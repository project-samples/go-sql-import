package main

import (
	"context"
	"fmt"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	cfg.Sql.Driver = "postgres"
	cfg.Sql.DataSourceName = "postgres://postgres:abcd1234@localhost/masterdata?sslmode=disable"

	ctx := context.Background()
	fmt.Println("Import file")
	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		fmt.Println("Error when initialize: ", err.Error())
		panic(err)
	}

	total, success, err := app.Import(ctx)
	if err != nil {
		fmt.Println("Error when import: ", err.Error())
		panic(err)
	}
	fmt.Println(fmt.Sprintf("total: %d, success: %d", total, success))
	fmt.Println(ctx, "Imported file")
}
