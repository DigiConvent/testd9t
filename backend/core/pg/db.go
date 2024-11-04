package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDatabaseInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

type PostgresDatabase struct {
	Conn *pgxpool.Pool
}

var DB PostgresDatabaseInterface

func GetPGDBConnection() PostgresDatabaseInterface {
	if DB == nil {
		username := os.Getenv("DB_USER")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		password := os.Getenv("DB_PASSWORD")
		dbName := "digiconvent"

		conn, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		DB = &PostgresDatabase{Conn: conn}
	}
	return DB
}

func (db *PostgresDatabase) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	row := db.Conn.QueryRow(ctx, sql, args...)
	return row
}

func (db *PostgresDatabase) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	rows, err := db.Conn.Query(ctx, sql, args...)
	return rows, err
}

func (db *PostgresDatabase) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	tag, err := db.Conn.Exec(ctx, sql, args...)
	return tag, err
}

func (db *PostgresDatabase) Close() {
	db.Conn.Close()
}

func (db *PostgresDatabase) RawSQL(sql string) string {
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal(err)
		}

		rowMap := make(map[string]interface{})
		for i, col := range rows.FieldDescriptions() {
			rowMap[string(col.Name)] = values[i]
		}

		results = append(results, rowMap)
	}

	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}
