package adapter

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"template/config"
	"template/constant"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewDatabase(cfg config.DatabaseConfig) (*sql.DB, error) {
	var (
		dbURL string
		err   error
	)
	once.Do(func() {
		switch cfg.Driver {
		case constant.DatabaseMysql:
			dbURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		case constant.DatabasePostgres:
			dbURL = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name)
		}

		db, err = sql.Open(cfg.Driver, dbURL)
		if err != nil {
			return
		}

		if cfg.ActivePool {
			db.SetMaxIdleConns(cfg.MaxPool) // Maximum number of idle connections
			db.SetMaxOpenConns(cfg.MinPool) // Maximum number of open connections
		}

		// Test the database connection
		err = db.Ping()
	})

	return db, err
}
