package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/242617/pace/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() error {

	var err error
	db, err = sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
