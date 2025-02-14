package post_service

import (
	"fmt"
	"net"
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

	conn, err := net.Dial("tcp", mx[0].Host+":25")
	if err != nil {
		return core.InternalError(err.Error())
	} else {
		log.Info("Connected to MX host: " + mx[0].Host)
	}
	defer conn.Close()

	domain := os.Getenv(constants.DOMAIN)

	fmt.Fprintln(conn, "HELO "+domain)
	fmt.Fprintln(conn, "MAIL FROM:<"+sender.Name+"@"+domain+">")
	fmt.Fprintln(conn, "RCPT TO:<"+to+">")
	fmt.Fprintln(conn, "DATA")
	fmt.Fprintln(conn, "Subject: "+subject+"\n\n"+body)
	fmt.Fprintln(conn, ".")
	fmt.Fprintln(conn, "QUIT")

	return core.StatusSuccess()
}
