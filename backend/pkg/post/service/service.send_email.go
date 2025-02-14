package post_service

import (
	"crypto/tls"
	"net/smtp"
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
	"github.com/google/uuid"
)

func (s PostService) SendEmail(from *uuid.UUID, to string, subject string, body string) *core.Status {
	if from == nil {
		return core.UnprocessableContentError("PostService requires an ID")
	}

	sender, status := s.ReadEmailAddress(from)
	if status.Err() {
		return status
	}

	senderEmail := sender.Name + "@" + sender.Domain

	addr := sender.Domain + ":" + os.Getenv(constants.SMTP_PORT)
	msg := "Subject: " + subject + "\r\n" +
		"From: " + senderEmail + "\r\n" +
		"To: " + to + "\r\n" +
		"\r\n" +
		body + "\r\n"

	auth := smtp.PlainAuth("", senderEmail, os.Getenv(constants.MASTER_PASSWORD), sender.Domain)

	go func() {
		err := smtp.SendMail(addr, auth, to, []string{to}, []byte(msg))
		if err != nil {
			log.Error("Unable to send electronic mail: " + err.Error())
		}
		log.Success("Email sent from " + senderEmail + " to " + to)
	}()
	go func() {
		auth := smtp.PlainAuth("", senderEmail, os.Getenv(constants.MASTER_PASSWORD), sender.Domain)

		keypair, err := tls.LoadX509KeyPair(post_setup.TlsPublicKeyPath(), post_setup.TlsPrivateKeyPath())
		if err != nil {
			log.Error("Unable to load x509 key pair: " + err.Error())
			return
		}
		log.Info("Loaded x509 key pair")

		tlsConfig := tls.Config{Certificates: []tls.Certificate{keypair}, ServerName: sender.Domain}

		conn, err := tls.Dial("tcp", addr, &tlsConfig)
		if err != nil {
			log.Error("Unable to connect to smtp server: " + err.Error())
			return
		}
		log.Info("Connected to smtp server")

		client, err := smtp.NewClient(conn, sender.Domain)
		if err != nil {
			log.Error("Unable to create smtp client: " + err.Error())
			return
		}
		log.Info("Created smtp client")

		if err = client.StartTLS(&tlsConfig); err != nil {
			log.Error("Unable to start TLS: " + err.Error())
			return
		}
		log.Info("Upgraded connection to TLS")

		if err = client.Auth(auth); err != nil {
			log.Error("Unable to authenticate: " + err.Error())
			return
		}
		log.Info("Authenticated with smtp server")

		if err = client.Mail(senderEmail); err != nil {
			log.Error("Unable to send mail: " + err.Error())
			return
		}
		log.Info("Sent mail from " + senderEmail)

		if err = client.Rcpt(to); err != nil {
			log.Error("Unable to set recipient: " + err.Error())
			return
		}
		log.Info("Set recipient to " + to)

		w, err := client.Data()
		if err != nil {
			log.Error("Unable to get data writer: " + err.Error())
			return
		}
		log.Info("Got data writer")

		_, err = w.Write([]byte(msg))
		if err != nil {
			log.Error("Unable to write data: " + err.Error())
			return
		}
		log.Info("Wrote data")

		if err = w.Close(); err != nil {
			log.Error("Unable to close data writer: " + err.Error())
			return
		}
		log.Info("Closed data writer")

		if err = client.Quit(); err != nil {
			log.Error("Unable to close smtp client: " + err.Error())
			return
		}
		log.Info("Closed smtp client")

		log.Success("Email sent from " + senderEmail + " to " + to)
	}()
	return core.IsProcessing()
}
