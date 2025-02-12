package post_setup

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path"
	"time"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
)

func TlsPrivateKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "", "privkey.pem")
}
func TlsPublicKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "", "fullchain.pem")
}
func TlsCaCertificatePath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "", "cacert.pem")
}
func TlsCaPrivateKeyPath() string {
	return path.Join(os.Getenv(constants.CERTIFICATES_PATH), "", "capk.pem")
}

func Setup() {
	log.Info("Executing setup for post")
	// These are not generated when deployed. That is done by letsencrypt
	if _, err := os.Stat(TlsPrivateKeyPath()); os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(TlsPrivateKeyPath()), 0755)
		if err != nil {
			panic("Cannot create folders for jwt: " + err.Error())
		}
		caCert, privateKey := getOrCreateCaCert()

		cert := &x509.Certificate{
			SerialNumber: big.NewInt(0),
			Subject: pkix.Name{
				CommonName:   os.Getenv("DOMAIN"),
				Organization: []string{os.Getenv("DOMAIN")},
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(10, 0, 0),
			KeyUsage:              x509.KeyUsageDigitalSignature,
			BasicConstraintsValid: true,
		}
		certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, &privateKey.PublicKey, privateKey)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(TlsPublicKeyPath(), pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes}), 0644)

		if err != nil {
			panic("Cannot create public key for jwt: " + err.Error())
		}
		// create privkey.pem
		err = os.WriteFile(TlsPrivateKeyPath(), pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}), 0644)
		if err != nil {
			panic("Cannot create private key for jwt: " + err.Error())
		}
	} else {
		log.Info("File exists: " + TlsPrivateKeyPath())
	}
}

func getOrCreateCaCert() (*x509.Certificate, *rsa.PrivateKey) {
	if _, err := os.Stat(TlsCaCertificatePath()); err == nil {
		os.MkdirAll(path.Dir(TlsCaPrivateKeyPath()), 0755)
		cert, err := os.ReadFile(TlsCaCertificatePath())
		if err != nil {
			panic(err)
		}
		block, _ := pem.Decode(cert)
		if block == nil {
			panic("Could not decode ca cert")
		}
		parsed, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			panic(err)
		}
		key, err := os.ReadFile(TlsCaPrivateKeyPath())
		if err != nil {
			panic(err)
		}
		block, _ = pem.Decode(key)
		if block == nil {
			panic("Could not decode ca key")
		}
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			panic(err)
		}
		return parsed, privateKey
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil
	}

	certTemplate := x509.Certificate{
		SerialNumber: big.NewInt(0),
		Subject: pkix.Name{
			CommonName:   os.Getenv("DOMAIN"),
			Organization: []string{os.Getenv("DOMAIN")},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{os.Getenv("DOMAIN")},
	}

	caCert, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil
	}

	parsed, err := x509.ParseCertificate(caCert)
	if err != nil {
		return nil, nil
	}

	pemCert := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCert,
	}
	err = os.WriteFile(TlsCaCertificatePath(), pem.EncodeToMemory(&pemCert), 0644)
	if err != nil {
		return nil, nil
	}

	pemPk := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	err = os.WriteFile(TlsCaPrivateKeyPath(), pem.EncodeToMemory(&pemPk), 0644)
	if err != nil {
		return nil, nil
	}

	return parsed, privateKey
}
