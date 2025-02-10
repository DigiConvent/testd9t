package post_service

import (
	"crypto/tls"
	"log"
	"net/smtp"

	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s PostService) SendEmail(from *uuid.UUID, to string, subject string, body string) *core.Status {
	fromEmail, status := s.Repository.ReadEmailAddress(from)
	if status.Err() {
		return &status
	}

	smtpServer := fromEmail.Domain
	port := "465"
	email := fromEmail.Name + "@" + fromEmail.Domain

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpServer,
	}

	conn, _ := tls.Dial("tcp", smtpServer+":"+port, tlsConfig)
	client, _ := smtp.NewClient(conn, smtpServer)
	defer client.Quit()
	auth := smtp.PlainAuth("", email, "", smtpServer)
	if err := client.Auth(auth); err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	message := []byte("Subject: Test Email\r\n" +
		"\r\n" +
		"This is a test email sent from a Go program.\r\n")

	if err := smtp.SendMail(smtpServer+":"+port, auth, email, []string{to}, message); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	return nil
}
