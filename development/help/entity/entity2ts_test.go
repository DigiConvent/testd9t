package entity_test

import (
	entity "development/help/entity"
	"fmt"
	"testing"
)

var testEntity string = "type User struct {\n" +
	"    ID        uuid.UUID   `json:\"id\"`\n" +
	"    ParentId  *uuid.UUID  `json:\"parent_id\"`\n" +
	"    Name      string      `json:\"name\"`\n" +
	"    Names     []string    `json:\"name\"`\n" +
	"    Amount    float64     `json:\"amount\"`\n" +
	"    IsActive  bool        `json:\"is_active\"`\n" +
	"    Timestamp *time.Time  `json:\"timestamp\"`\n" +
	"    Data      []byte      `json:\"data\"`\n" +
	"}"

var expectedEntity string = "export interface User {\n" +
	"id string\n" +
	"parent_id string | null\n" +
	"name string\n" +
	"names string[]\n" +
	"amount number\n" +
	"is_active bool\n" +
	"timestamp Date\n" +
	"data string\n" +
	"}"

func TestGolangStructToTypeScriptType(t *testing.T) {
	fmt.Println(testEntity)
	var ts string = entity.GolangStructToTypeScriptType(testEntity)
	t.Log("\n" + ts)
	if ts != expectedEntity {
		t.Error("Failed to convert Golang struct to TypeScript")
	}
	// t.Error(lineToTest)
	// t.Log(entity.GoStructLineToTypeScript(lineToTest))
}

func TestCamelCase(t *testing.T) {
	cases := []string{
		"ID",
		"UserId",
		"FirstName",
		"lastName",
		"DATETIME",
		"NotSureWhereThisBelongs",
	}

	expectation := []string{
		"id",
		"user_id",
		"first_name",
		"last_name",
		"",
		"not_sure_where_this_belongs",
	}

	for i := range cases {
		result, _ := entity.PascalToCamelCase(cases[i])
		if result != expectation[i] {
			t.Error("Failed to convert", cases[i], "to", expectation[i], " ("+result+")")
		} else {
			t.Log("Converted", cases[i], "to", result)
		}
	}
}
