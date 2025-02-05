package db_test

import (
	"os"
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
)

func TestSqliteDB(t *testing.T) {
	os.Stdout = nil
	testDB := db.NewTestSqliteDB("core.db.sqlite")
	defer testDB.DeleteDatabase()

	res, err := testDB.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}

	res, err = testDB.Exec("INSERT INTO test (name) VALUES (?)", "testthis")

	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}

	rows, err := testDB.Query("SELECT name FROM test")

	if err != nil {
		t.Fatal(err)
	}

	if rows == nil {
		t.Fatal("Expected a result")
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			t.Fatal(err)
		}

		if name != "testthis" {
			t.Fatalf("Expected 'testthis', got '%s'", name)
		}
	}
}
