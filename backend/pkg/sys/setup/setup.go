package sys_setup

import (
	"path"

	constants "github.com/DigiConvent/testd9t/core/const"
)

func TlsPrivateKeyPath() string {
	return path.Join(constants.CERTIFICATES_PATH, "privkey.pem")
}
func TlsPublicKeyPath() string {
	return path.Join(constants.CERTIFICATES_PATH, "fullchain.pem")
}
