package db

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDBConnect(config Config) (*sqlx.DB, error) {
	cfg := mysql.Config{
		User:   config.Username,
		Passwd: config.Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", config.Host, config.Port),
		DBName: config.DBName,
	}
	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
