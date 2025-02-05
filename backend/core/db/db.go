package db

import "database/sql"

type DatabaseInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	pkgDir() string
	MigratePackage() error
	DeleteDatabase()
	Close()
}
