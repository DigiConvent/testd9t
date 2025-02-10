package iam_service_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path"
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
)

var testDB db.DatabaseInterface

const privKeyPath = "/tmp/testd9t/test/privkey.pem"

func GetTestIAMService(dbName string) iam_service.IAMServiceInterface {
	if testDB == nil {
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := iam_repository.NewIAMRepository(testDB, privKeyPath, false)
	return iam_service.NewIAMService(repo)
}

func TestMain(m *testing.M) {
	generatePrivateKey()
	GetTestIAMService("iam")
	defer testDB.DeleteDatabase()
	m.Run()
}

func generatePrivateKey() {
	if _, err := os.Stat(privKeyPath); err == nil {
		return
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	err = privateKey.Validate()
	if err != nil {
		panic(err)
	}

	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	folder := path.Dir(privKeyPath)
	os.MkdirAll(folder, 0755)
	privatePEM := pem.EncodeToMemory(&privBlock)
	os.WriteFile(privKeyPath, privatePEM, 0644)
}
