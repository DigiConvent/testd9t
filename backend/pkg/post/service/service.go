package post_service

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
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
		go s.handleSMTPConnection(connection)
	}
}

func (s *PostService) handleSMTPConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "220 SMTP Server for testd9t is ready")

	scanner := bufio.NewScanner(conn)

	authenticated := false
	isData := false
	var sender, recipient, data string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Received:", line)

		if strings.HasPrefix(line, "HELO") || strings.HasPrefix(line, "EHLO") {
			fmt.Fprintln(conn, "250-Hello")
			fmt.Fprintln(conn, "250 AUTH PLAIN")
		} else if strings.HasPrefix(line, "AUTH PLAIN") {
			fmt.Fprintln(conn, "334")

			scanner.Scan()
			decoded, _ := base64.StdEncoding.DecodeString(scanner.Text())
			parts := strings.SplitN(string(decoded), "\x00", 3)
			email := parts[1]
			password := parts[2]

			_, status := s.repository.GetEmailAddressByName(email)
			if status.Err() {
				fmt.Fprintln(conn, "535 Authentication failed")
				continue
			}

			if len(parts) == 3 && password == os.Getenv("MASTER_PASSWORD") {
				authenticated = true
				fmt.Fprintln(conn, "235 Authentication successful")
			} else {
				fmt.Fprintln(conn, "535 Authentication failed")
			}
		} else if strings.HasPrefix(line, "MAIL FROM:") {
			if !authenticated {
				fmt.Fprintln(conn, "530 Authentication required")
				continue
			}
			sender = strings.TrimPrefix(line, "MAIL FROM:")
			fmt.Fprintln(conn, "250 OK")
		} else if strings.HasPrefix(line, "RCPT TO:") {
			if !authenticated {
				fmt.Fprintln(conn, "530 Authentication required")
				continue
			}
			recipient = strings.TrimPrefix(line, "RCPT TO:")
			fmt.Fprintln(conn, "250 OK")
		} else if strings.HasPrefix(line, "DATA") {
			if !authenticated {
				fmt.Fprintln(conn, "530 Authentication required")
				continue
			}
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
