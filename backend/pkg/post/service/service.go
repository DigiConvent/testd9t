package post_service

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	"github.com/google/uuid"
)

type PostServiceInterface interface {
	CreateEmailAddress(credentials *post_domain.EmailAddressWrite) (*uuid.UUID, *core.Status)
	ReadEmailAddress(id *uuid.UUID) (*post_domain.EmailAddressRead, *core.Status)
	DeleteEmailAddress(id *uuid.UUID) *core.Status
	ListEmailAddresses() ([]post_domain.EmailAddressRead, *core.Status)
	UpdateEmailAddresses(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) *core.Status

	SendEmail(from *uuid.UUID, to, subject, body string) *core.Status
}

type PostService struct {
	repository post_repository.PostRepositoryInterface
	address    string
}

func NewPostService(repository post_repository.PostRepositoryInterface, live bool) PostServiceInterface {
	postService := PostService{
		repository: repository,
		address:    ":2525",
	}
	if live {
		postService.StartSmtpServer()
	}
	return postService
}

func (s *PostService) StartSmtpServer() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		panic(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Error("Error accepting connection: " + err.Error())
			continue
		}
		go s.startListening(connection)
	}
}

func (s PostService) startListening(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "220 GoSMTP Server Ready")

	scanner := bufio.NewScanner(conn)
	var sender, recipient, data string
	isData := false

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Client:", line)

		if strings.HasPrefix(line, "HELO") {
			fmt.Fprintln(conn, "250 Hello")
		} else if strings.HasPrefix(line, "MAIL FROM:") {
			sender = strings.TrimPrefix(line, "MAIL FROM:")
			fmt.Fprintln(conn, "250 OK")
		} else if strings.HasPrefix(line, "RCPT TO:") {
			recipient = strings.TrimPrefix(line, "RCPT TO:")
			fmt.Fprintln(conn, "250 OK")
		} else if strings.HasPrefix(line, "DATA") {
			fmt.Fprintln(conn, "354 End data with <CR><LF>.<CR><LF>")
			isData = true
			data = ""
		} else if isData {
			if line == "." {
				fmt.Println("Received Email:")
				fmt.Println("From:", sender)
				fmt.Println("To:", recipient)
				fmt.Println("Data:", data)
				fmt.Fprintln(conn, "250 Message accepted for delivery")
				isData = false
			} else {
				data += line + "\n"
			}
		} else if strings.HasPrefix(line, "QUIT") {
			fmt.Fprintln(conn, "221 Bye")
			break
		} else {
			fmt.Fprintln(conn, "500 Unrecognized command")
		}
	}
}
