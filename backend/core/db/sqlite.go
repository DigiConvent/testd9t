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
var DatabasePath string = "/tmp/testd9t/db/"

func CloseAllDatabases() {
	for _, db := range databases {
		db.Close()
	}
}

func NewTestSqliteDB(dbName string) DatabaseInterface {
	if databases[dbName] != nil {
		databases[dbName].Close()
		delete(databases, dbName)
	}
	connection := SqliteConnection(dbName, true)

	if connection.pkgDir() != "" {
		err := connection.MigratePackage()
		if err != nil {
			connection.DeleteDatabase()
			panic(err)
		}
	}

	return connection
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
		if _, err := os.Stat(dbPath); err == nil {
			log.Success("Loading existing database at " + dbPath)
		}

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

func packageName() string {
	thisPath, _ := os.Getwd()
	segments := strings.Split(thisPath, "/testd9t/backend/pkg/")
	if len(segments) < 2 {
		return ""
	}
	packageName := strings.Split(segments[1], "/")[0]
	return packageName
}

func (s *SqliteDatabase) pkgDir() string {
	// project dir
	workingDir, _ := os.Getwd()
	segments := strings.Split(workingDir, "/testd9t/")
	if len(segments) < 2 {
		return ""
	}
	return path.Join(segments[0], "testd9t", "backend", "pkg", s.name)
}

func (s *SqliteDatabase) MigratePackage() error {
	pkgName := packageName()
	if pkgName == "" {
		return nil
	}
	dbPath := path.Join(s.pkgDir(), "db")

	versions, err := os.ReadDir(dbPath)

	if err != nil {
		return err
	}

	for _, version := range versions {
		if version.IsDir() {
			migrations, err := os.ReadDir(path.Join(dbPath, version.Name()))

			if err != nil {
				return err
			}

			for _, migration := range migrations {
				sql, err := os.ReadFile(path.Join(dbPath, version.Name(), migration.Name()))
				if err != nil {
					return err
				}

				result, err := s.DB.Exec(string(sql))

				if err != nil {
					log.Error("❌ " + s.name + ":" + migration.Name())
					return err
				}

				if result == nil {
					log.Error("Migration did not return a result: " + migration.Name() + " on database " + s.name)
					return nil
				}
			}
		}
	}

	return nil
}

func (s *SqliteDatabase) Dir() string {
	return path.Join(DatabasePath, s.name)
}

func (s *SqliteDatabase) DeleteDatabase() {
	s.DB.Close()
	var err error
	if s.test {
		err = os.RemoveAll(s.Dir())
	} else {
		err = os.Remove(s.Dir())
	}
	if err != nil {
		log.Error("Could not delete database: " + s.Dir())
		panic(err)
	}

	if _, err := os.Stat(s.Dir()); err == nil {
		log.Error("Database still exists: " + s.Dir())
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
