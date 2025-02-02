package db

import (
	"database/sql"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/DigiConvent/testd9t/core/log"
	_ "github.com/mattn/go-sqlite3"
)

var databases = map[string]DatabaseInterface{}
var DatabasePath string = "/tmp/d9t/db/"

func CloseAllDatabases() {
	for _, db := range databases {
		db.Close()
	}
}

func NewTestSqliteDB(dbName string) DatabaseInterface {
	return SqliteConnection(dbName, true)
}

func NewSqliteDB(dbName string) DatabaseInterface {
	return SqliteConnection(dbName, false)
}

func SqliteConnection(dbName string, test bool) DatabaseInterface {
	dbName = strings.ToLower(dbName)
	is_alphanumeric := regexp.MustCompile(`^[a-zA-Z\.]*$`).MatchString(dbName)
	if !is_alphanumeric {
		panic("Database name must be alphanumeric")
	}

	dbPath := path.Join(DatabasePath, dbName)

	if databases[dbName] == nil {
		var db *sql.DB
		var err error
		err = os.MkdirAll(dbPath, 0755)

		if err != nil {
			log.Error("Could not create database directory: " + dbPath)
			panic(err)
		}

		dbPath = path.Join(dbPath, "database.db")
		db, err = sql.Open("sqlite3", dbPath)

		if err != nil {
			log.Error("Could not create/open database: " + dbPath)
			panic(err)
		}

		databases[dbName] = &SqliteDatabase{
			DB:   db,
			name: dbName,
			test: test,
		}
	}

	return databases[dbName]
}

type SqliteDatabase struct {
	DB   *sql.DB
	name string
	test bool
}

func ListPackages() []string {
	var packages []string

	for key := range databases {
		packages = append(packages, key)
	}

	return packages
}

func packagePath() string {
	thisPath, _ := os.Getwd()
	segments := strings.Split(thisPath, "/pkg/")
	packagePath := segments[0] + "/pkg"
	packageName := strings.Split(segments[1], "/")[0]
	dbPath := path.Join(packagePath, packageName)
	return dbPath
}

func (s *SqliteDatabase) MigratePackage() {
	dbPath := path.Join(packagePath(), "db")

	dir, err := os.ReadDir(dbPath)

	if err != nil {
		panic(err)
	}

	for _, version := range dir {
		if version.IsDir() {
			migrations, err := os.ReadDir(path.Join(dbPath, version.Name()))

			if err != nil {
				panic(err)
			}

			for _, migration := range migrations {
				sql, err := os.ReadFile(path.Join(dbPath, version.Name(), migration.Name()))
				if err != nil {
					panic(err)
				}

				result, err := s.DB.Exec(string(sql))

				if err != nil {
					panic(string(sql) + err.Error())
				}

				if result == nil {
					panic("Expected a result")
				}
			}
		}
	}
}

func (s *SqliteDatabase) DeleteDatabase() {
	s.DB.Close()
	dbPath := path.Join(DatabasePath, s.name)
	var err error
	if s.test {
		err = os.Remove(path.Join("/tmp", s.name+".db"))
	} else {
		err = os.Remove(dbPath)
	}
	if err != nil {
		panic(err)
	}
}

func (s *SqliteDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *SqliteDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.Query(query, args...)
}

func (s *SqliteDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.DB.QueryRow(query, args...)
}

func (s *SqliteDatabase) Close() {
	s.DB.Close()
}
