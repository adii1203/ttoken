package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func connectDb() (*pgx.Conn, error) {

	dns := "postgresql://postgres.refgvcbesudvjixvhqit:IAJib3bpxETDkR3l@aws-0-ap-south-1.pooler.supabase.com:5432/postgres"
	db, err := pgx.Connect(context.Background(), dns)

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
