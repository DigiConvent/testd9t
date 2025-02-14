// exempt from testing
package post_service

import (
	"net"
	"net/smtp"
	"os"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	"github.com/google/uuid"
)

func (s PostService) SendEmail(from *uuid.UUID, to string, subject string, body string) *core.Status {
	if from == nil {
		return core.UnprocessableContentError("PostService requires an ID")
	}

	sender, status := s.ReadEmailAddress(from)
	if status.Err() {
		return status
	} else {
		log.Info("Sender: " + sender.Name)
	}

	mx, err := net.LookupMX(strings.Split(to, "@")[1])
	if err != nil {
		return core.InternalError(err.Error())
	} else {
		log.Info("Found MX records: " + mx[0].Host)
	}

	client, err := smtp.Dial(mx[0].Host + ":25")
	if err != nil {
		return core.InternalError(err.Error())
	} else {
		log.Info("Connected to MX host: " + mx[0].Host)
	}
	defer client.Close()

	domain := os.Getenv(constants.DOMAIN)

	client.Hello(domain)

	return core.StatusSuccess()
}
