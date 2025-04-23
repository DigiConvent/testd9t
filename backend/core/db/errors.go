package db

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
)

func SqliteError(err error) *core.Status {
	if err == nil {
		return nil
	}

	if strings.HasPrefix(err.Error(), "sql: no rows in result set") {
		return core.NotFoundError(err.Error())
	}
	return nil
}
