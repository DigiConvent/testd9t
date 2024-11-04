package db_test

import (
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestPostgresDatabase(t *testing.T) {
	// test postgres database
	pool, err := pgxmock.NewPool()
	if err != nil {
		panic(err)
	}
	if pool == nil {
		t.Error("expected db not to be nil")
	}
}
