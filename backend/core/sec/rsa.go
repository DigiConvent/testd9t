package sec

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKeyPair(bits int) (privatePEM string, publicPEM string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePEMBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateBytes,
	}
	privatePEM = string(pem.EncodeToMemory(privatePEMBlock))

	publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicPEMBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicBytes,
	}
	publicPEM = string(pem.EncodeToMemory(publicPEMBlock))

	return privatePEM, publicPEM, nil
}
