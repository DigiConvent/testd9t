package entity

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

type Field struct {
	Name string
	Type string
}

func GolangStructToTypeScriptType(input string) string {
	lines := strings.Split(input, "\n")
	structName := ""
	var fields []Field = make([]Field, 0)
	for _, line := range lines {
		if strings.Contains(line, "type") && strings.Contains(line, "struct {") {
			structName = strings.Split(line, " ")[1]
		} else {
			if line == "" || strings.HasPrefix(line, "}") || strings.HasPrefix(line, "//") {
				continue
			}
			lineSegments := strings.Fields(line)
			fields = append(fields, Field{Name: lineSegments[0], Type: lineSegments[1]})
		}
	}

	typescriptType := "export interface " + structName + " {\n"
	for _, field := range fields {
		fieldName, err := PascalToCamelCase(field.Name)
		if err != nil {
			panic(err)
		}
		typescriptType += ("\t" + fieldName + ": " + GoToTypeScript(field.Type) + ";\n")
	}
	typescriptType += "}\n"

	return typescriptType
}

func GoStructLineToTypeScript(line string) string {
	line = strings.TrimSpace(line)
	segments := strings.Split(line, " ")
	if len(segments) < 2 {
		panic("Invalid line")
	}
	return segments[0] + " " + GoToTypeScript(segments[1])
}

func GoToTypeScript(goType string) string {
	isOptional := strings.HasPrefix(goType, "*")
	goType = strings.TrimPrefix(goType, "*")
	isArray := strings.HasPrefix(goType, "[]")
	goType = strings.TrimPrefix(goType, "[]")
	var tsType string

	switch goType {
	case "int", "int32", "int64":
		tsType = "number"
	case "uuid.UUID":
		tsType = "string"
	case "float32", "float64":
		tsType = "number"
	case "bool":
		tsType = "boolean"
	case "string":
		tsType = "string"
	case "byte":
		if isArray {
			tsType = "Uint8Array"
		}
	case "interface{}":
		tsType = "any"
	case "time.Time":
		tsType = "Date"
	default:
		tsType = "any (" + goType + ")"
	}

	if isArray {
		tsType += "[]"
	}
	if isOptional {
		tsType += " | null"
	}

	return tsType
}

func PascalToCamelCase(input string) (string, error) {
	input = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(input, "")
	if input == "ID" {
		return "id", nil
	}

	formattedWord := ""
	for i, c := range input {
		if i > 0 && unicode.IsUpper(c) && unicode.IsUpper(rune(input[i-1])) {
			return "", errors.New("Illegal: " + input)
		}
		if unicode.IsUpper(c) {
			formattedWord += "_" + strings.ToLower(string(c))
		} else {
			formattedWord += string(c)
		}
	}

	formattedWord, _ = strings.CutPrefix(formattedWord, "_")
	return formattedWord, nil
	// newWord := ""
	// for i, c := range strings.Split(formattedWord, "_") {
	// 	if i != 0 {
	// 		newWord += strings.ToUpper(string(c[0])) + c[1:]
	// 	} else {
	// 		newWord += c
	// 	}
	// }

	// return newWord, nil
}
