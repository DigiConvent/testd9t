package post_service

import (
	"fmt"
	"net/smtp"

	"github.com/DigiConvent/testd9t/core"
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

	addr := "localhost:2525"
	msg := "Subject: " + subject + "\r\n" +
		"From: " + senderEmail + "\r\n" +
		"To: " + to + "\r\n" +
		"\r\n" +
		body + "\r\n"

	err := smtp.SendMail(addr, nil, to, []string{to}, []byte(msg))
	if err != nil {
		return core.InternalError(err.Error())
	}
	fmt.Println("Email sent successfully!")
	return core.StatusSuccess()
}
