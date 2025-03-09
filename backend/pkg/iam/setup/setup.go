package iam_setup

import (
	"os"
	"path"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	"github.com/DigiConvent/testd9t/core/sec"
)

func JwtPublicKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "iam", "pubkey.pem")
}
func JwtPrivateKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "iam", "privkey.pem")
}
func Setup() {
	log.Info("Executing setup for iam")
	if _, err := os.Stat(JwtPublicKeyPath()); os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(JwtPublicKeyPath()), 0755)
		if err != nil {
			log.Error("Cannot create folders for jwt: " + err.Error())
		}

		privatePem, publicPem, _ := sec.GenerateRSAKeyPair(1024)

		err = os.WriteFile(JwtPrivateKeyPath(), []byte(privatePem), 0644)
		if err != nil {
			log.Error("Cannot create private key for jwt: " + err.Error())
		}

		err = os.WriteFile(JwtPublicKeyPath(), []byte(publicPem), 0644)
		if err != nil {
			log.Error("Cannot create public key for jwt: " + err.Error())
		}
	}
}
