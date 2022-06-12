package app

import "github.com/core-go/sql"

type Config struct {
	Sql    sql.Config         `mapstructure:"sql"`
}
