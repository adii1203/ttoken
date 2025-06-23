package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func connectDb() (*pgx.Conn, error) {
	dns := ""
	config, err := pgx.ParseConfig(dns)
	if err != nil {
		return nil, err
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB() *pgx.Conn {
	db, err := connectDb()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
