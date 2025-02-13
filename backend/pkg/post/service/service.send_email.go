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

	log.Info("Creating auth object")
	auth := smtp.PlainAuth("", senderEmail, os.Getenv(constants.MASTER_PASSWORD), sender.Domain)
	log.Info("Created  auth object")

	go func() {
		err := smtp.SendMail(addr, auth, to, []string{to}, []byte(msg))
		if err != nil {
			log.Error("Unable to send electronic mail: " + err.Error())
		}
		log.Success("Email sent from " + senderEmail + " to " + to)
	}()
	go func() {
		log.Info(1)
		auth := smtp.PlainAuth("", senderEmail, os.Getenv(constants.MASTER_PASSWORD), sender.Domain)
		log.Info(2)
		keypair, err := tls.LoadX509KeyPair(post_setup.TlsPublicKeyPath(), post_setup.TlsPrivateKeyPath())
		log.Info(3)
		if err != nil {
			log.Error("Unable to load x509 key pair: " + err.Error())
		}
		log.Info(4)
		tlsConfig := tls.Config{Certificates: []tls.Certificate{keypair}}
		log.Info(5)
		conn, err := tls.Dial("tcp", addr, &tlsConfig)
		if err != nil {
			log.Error("Unable to connect to smtp server: " + err.Error())
		}

		log.Info(6)
		client, err := smtp.NewClient(conn, sender.Domain)
		if err != nil {
			log.Error("Unable to create smtp client: " + err.Error())
		} else {
			log.Info("Created smtp client")
		}

		err = client.Auth(auth)
		if err != nil {
			log.Error("Unable to authenticate: " + err.Error())
		} else {
			log.Info("Authenticated with smtp server")
		}

		err = client.Mail(senderEmail)
		if err != nil {
			log.Error("Unable to send mail: " + err.Error())
		} else {
			log.Info("Sent mail from " + senderEmail)
		}

		err = client.Rcpt(to)
		if err != nil {
			log.Error("Unable to set recipient: " + err.Error())
		} else {
			log.Info("Set recipient to " + to)
		}

		w, err := client.Data()
		if err != nil {
			log.Error("Unable to get data writer: " + err.Error())
		} else {
			log.Info("Got data writer")
		}

		_, err = w.Write([]byte(msg))
		if err != nil {
			log.Error("Unable to write data: " + err.Error())
		} else {
			log.Info("Wrote data")
		}

		err = w.Close()
		if err != nil {
			log.Error("Unable to close data writer: " + err.Error())
		} else {
			log.Info("Closed data writer")
		}

		err = client.Quit()
		if err != nil {
			log.Error("Unable to close smtp client: " + err.Error())
		} else {
			log.Info("Closed smtp client")
		}

		log.Success("Email sent from " + senderEmail + " to " + to)
	}()
	return core.IsProcessing()
}
