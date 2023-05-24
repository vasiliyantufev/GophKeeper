package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/config/configserver"
)

type DB struct {
	Pool *sql.DB
	log  *logrus.Logger
}

func New(config *configserver.ConfigServer, log *logrus.Logger) (*DB, error) {
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