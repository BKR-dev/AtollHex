package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type DBEngine interface {
	Connect() error
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)
	Close() error
}

type PostgresEngine struct {
	conn *pgx.Conn
}

func NewDBEngine() DBEngine {
	return &PostgresEngine{}
}

func (db *PostgresEngine) Connect() error {
	var err error
	db.conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	return err
}

func (db *PostgresEngine) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, val := range values {
			row[rows.FieldDescriptions()[i].Name] = val
		}

		results = append(results, row)
	}

	return results, nil
}

func (db *PostgresEngine) Close() error {
	return db.conn.Close(context.Background())
}
