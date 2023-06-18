package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
)

type DB struct {
	Pool *sql.DB
	log  *logrus.Logger
}

func New(config *config.Config, log *logrus.Logger) (*DB, error) {
	pool, err := sql.Open("postgres", config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cnl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cnl()

	if err := pool.PingContext(ctx); err != nil {
		return nil, err
	}
	return &DB{Pool: pool}, nil
}

func (db DB) Close() error {
	return db.Pool.Close()
}

func (db DB) Ping() error {
	if err := db.Pool.Ping(); err != nil {
		db.log.Error(err)
		return err
	}
	return nil
}
