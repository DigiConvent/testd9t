package install

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DBHost       string `json:"DB_HOST"`
	DBPort       string `json:"DB_PORT"`
	DBUser       string `json:"DB_USER"`
	DBPassword   string `json:"DB_PASSWORD"`
	EnvPath      string `json:"ENV_PATH"`
	BinaryPath   string `json:"BINARY_PATH"`
	DataPath     string `json:"DATA_PATH"`
	FrontendPath string `json:"FRONTEND_PATH"`
}

var EnvironmentVars EnvVars = EnvVars{
	DBHost:       "localhost",
	DBPort:       "5432",
	DBUser:       "postgres",
	DBPassword:   "",
	EnvPath:      "/etc/digiconvent/env",
	BinaryPath:   "/usr/local/bin/digiconvent",
	DataPath:     "/var/digiconvent/",
	FrontendPath: "/var/digiconvent/frontend/",
}

func OverwriteFromFile(overwriteWithFile string, verbose bool) {
	godotenv.Load(overwriteWithFile)
	v := reflect.ValueOf(&EnvironmentVars).Elem()
	typeOfE := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := typeOfE.Field(i).Tag.Get("json")
		if key == "ENV_PATH" {
			continue
		}
		currentValue := os.Getenv(key)
		defaultValue := v.Field(i).String()
		v.Field(i).SetString(defaultValue)
		if currentValue != "" && currentValue != defaultValue && verbose {
			fmt.Println("Overwrote", key, "with", currentValue, " (previously", "'"+defaultValue+"')")
		}
	}
}

func (e *EnvVars) WriteEnvFile() {
	v := reflect.ValueOf(e)
	typeOfE := v.Type()
	contents := ""

	for i := 0; i < v.NumField(); i++ {
		key := typeOfE.Field(i).Tag.Get("json")
		value := v.Field(i).String()
		entry := key + "=" + value
		contents += entry + "\n"
	}

	os.WriteFile(e.EnvPath, []byte(contents), 0644)
}

/*
/etc/digiconvent/env
/etc/digiconvent/certs/
/etc/digiconvent/certs/fullchain.pem
/etc/digiconvent/certs/privkey.pem
/etc/digiconvent/migrations/
/etc/digiconvent/migrations/x.y.z/abc_.sql
*/
