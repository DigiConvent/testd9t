package env

import (
	"fmt"
	"os"
	"regexp"
)

func ReplaceEnvVar(input string) string {
	re := regexp.MustCompile(`\${([^}]*)}`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		varName := match[2 : len(match)-1]

		value, exists := os.LookupEnv(varName)
		if !exists {
			fmt.Printf("Environment variable %s not found\n", varName)
		}
		if !exists {
			return match
		}
		return value
	})
}
