package post_service

import (
	"net/smtp"
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
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

	log.Info(addr)
	log.Info(msg)

	auth := smtp.PlainAuth("", sender.Name, os.Getenv(constants.MASTER_PASSWORD), sender.Domain)

	err := smtp.SendMail(addr, auth, to, []string{to}, []byte(msg))
	if err != nil {
		log.Error("Unable to send electronic mail: " + err.Error())
		return core.InternalError("Unable to send electronic mail: " + err.Error())
	}
	log.Success("Email sent from " + senderEmail + " to " + to)
	return core.StatusSuccess()
}
