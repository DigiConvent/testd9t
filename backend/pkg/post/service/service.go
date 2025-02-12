package post_service

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
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

func NewPostService(repository post_repository.PostRepositoryInterface, startSmtpServer bool) PostServiceInterface {
	postService := PostService{
		repository: repository,
		address:    ":" + os.Getenv(constants.SMTP_PORT),
	}

	if startSmtpServer {
		log.Info("Starting smtp server on " + postService.address)
		go postService.StartSmtpServer()
	} else {
		log.Info("Skipping smtp server start")
	}
	return postService
}

func (s *PostService) StartSmtpServer() {
	cert, err := tls.LoadX509KeyPair(post_setup.TlsPublicKeyPath(), post_setup.TlsPrivateKeyPath())
	if err != nil {
		log.Error("Error loading certificate for the smtp server: " + err.Error())
	}
	listener, err := tls.Listen("tcp", s.address,
		&tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	)

	if err != nil {
		log.Error("Error starting smtp server: " + err.Error())
	}

	defer func() {
		log.Info("Closing smtp server (defer)")
		err := listener.Close()
		if err != nil {
			log.Error("Error closing smtp server: " + err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Info("Closing smtp server (signal)")
		err := listener.Close()
		if err != nil {
			log.Error("Error closing smtp server: " + err.Error())
		}
	}()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Error("Error accepting connection: " + err.Error())
			continue
		} else {
			log.Info("Accepted connection from " + connection.RemoteAddr().String())
		}
		go s.handleSMTPConnection(connection)
	}

}
func (s *PostService) handleSMTPConnection(conn net.Conn) {
	defer conn.Close()
	response := bufio.NewWriter(conn)
	fmt.Fprintln(response, "220 SMTP Server for testd9t is ready")
	response.Flush()

	scanner := bufio.NewScanner(conn)

	authenticated := false
	isData := false
	var sender, recipient, data string

	for scanner.Scan() {
		line := scanner.Text()
		log.Info("[SMTP] Received: " + line)

		if strings.HasPrefix(line, "HELO") || strings.HasPrefix(line, "EHLO") {
			fmt.Fprintln(response, "250-Hello")
			fmt.Fprintln(response, "250 AUTH PLAIN")
			response.Flush()
		} else if strings.HasPrefix(line, "AUTH PLAIN") {
			fmt.Fprintln(response, "334 ")
			response.Flush()

			if !scanner.Scan() {
				fmt.Fprintln(response, "535 Authentication failed")
				response.Flush()
				break
			}
			decoded, _ := base64.StdEncoding.DecodeString(scanner.Text())
			parts := strings.SplitN(string(decoded), "\x00", 3)

			if len(parts) < 3 {
				fmt.Fprintln(response, "535 Authentication failed")
				response.Flush()
				continue
			}

			email := parts[1]
			password := parts[2]

			_, status := s.repository.GetEmailAddressByName(email)
			if status.Err() {
				fmt.Fprintln(response, "535 Authentication failed")
				response.Flush()
				continue
			}

			if password == os.Getenv(constants.MASTER_PASSWORD) {
				authenticated = true
				fmt.Fprintln(response, "235 Authentication successful")
			} else {
				fmt.Fprintln(response, "535 Authentication failed")
			}
			response.Flush()
		} else if strings.HasPrefix(line, "MAIL FROM:") {
			if !authenticated {
				fmt.Fprintln(response, "530 Authentication required")
				response.Flush()
				continue
			}
			sender = strings.TrimSpace(strings.TrimPrefix(line, "MAIL FROM:"))
			fmt.Fprintln(response, "250 OK")
			response.Flush()
		} else if strings.HasPrefix(line, "RCPT TO:") {
			if !authenticated {
				fmt.Fprintln(response, "530 Authentication required")
				response.Flush()
				continue
			}
			recipient = strings.TrimSpace(strings.TrimPrefix(line, "RCPT TO:"))
			fmt.Fprintln(response, "250 OK")
			response.Flush()
		} else if strings.HasPrefix(line, "DATA") {
			if !authenticated {
				fmt.Fprintln(response, "530 Authentication required")
				response.Flush()
				continue
			}
			fmt.Fprintln(response, "354 End data with <CR><LF>.<CR><LF>")
			response.Flush()
			isData = true
			data = ""
		} else if isData {
			if strings.TrimSpace(line) == "." {
				fmt.Println("Received Email:")
				fmt.Println("From:", sender)
				fmt.Println("To:", recipient)
				fmt.Println("Data:", data)
				fmt.Fprintln(response, "250 OK: Message accepted")
				isData = false
			} else {
				data += line + "\n"
			}
		} else if strings.HasPrefix(line, "QUIT") {
			fmt.Fprintln(response, "221 Bye")
			response.Flush()
			break
		} else {
			fmt.Fprintln(response, "500 Unrecognized command")
			response.Flush()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("Connection error: " + err.Error())
	}
}
