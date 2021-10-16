package db

import (
	"context"
	"database/sql"
	"demo/blog/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dsn := config.GetConfig().GetString("dataSourceName")
	var err error
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	ping()

}

func ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
