package sys_setup

import (
	"os"
	"path"

	constants "github.com/DigiConvent/testd9t/core/const"
)

func TlsPrivateKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "privkey.pem")
}
func TlsPublicKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "fullchain.pem")
}
