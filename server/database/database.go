package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var instance *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dev123456"
	dbname   = "manager"
)

func Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	instance = db
	return nil
}

func NewTransaction(ctx context.Context) (*sql.Tx, error) {
	return instance.BeginTx(ctx, &sql.TxOptions{})
}
