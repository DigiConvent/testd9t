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

	log.Info("Sending email from " + from.String() + " to " + to)

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
		}
		tlsConfig := tls.Config{Certificates: []tls.Certificate{keypair}}
		conn, err := tls.Dial("tcp", addr, &tlsConfig)
		if err != nil {
			log.Error("Unable to connect to smtp server: " + err.Error())
		}

		client, err := smtp.NewClient(conn, sender.Domain)
		if err != nil {
			log.Error("Unable to create smtp client: " + err.Error())
		}

		err = client.Auth(auth)
		if err != nil {
			log.Error("Unable to authenticate: " + err.Error())
		}

		err = client.Mail(senderEmail)
		if err != nil {
			log.Error("Unable to send mail: " + err.Error())
		}

		err = client.Rcpt(to)
		if err != nil {
			log.Error("Unable to set recipient: " + err.Error())
		}

		w, err := client.Data()
		if err != nil {
			log.Error("Unable to get data writer: " + err.Error())
		}

		_, err = w.Write([]byte(msg))
		if err != nil {
			log.Error("Unable to write data: " + err.Error())
		}
		err = w.Close()
		if err != nil {
			log.Error("Unable to close data writer: " + err.Error())
		}
		err = client.Quit()
		if err != nil {
			log.Error("Unable to close smtp client: " + err.Error())
		}

		log.Success("Email sent from " + senderEmail + " to " + to)
	}()
	return core.IsProcessing()
}
