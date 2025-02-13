// exempt from testing
package post_service

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
)

func (s *PostService) startSmtpServer() {
	cert, err := tls.LoadX509KeyPair(post_setup.TlsPublicKeyPath(), post_setup.TlsPrivateKeyPath())
	if err != nil {
		log.Error("Error loading certificate for the smtp server: " + err.Error())
	}

	tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", s.address, &tlsConfig)
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
		go s.handleSmtpConnection(connection)
	}
}

func (s *PostService) handleSmtpConnection(conn net.Conn) {
	log.Info("[SMTP] Connection from " + conn.RemoteAddr().String())
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Error closing [SMTP] connection: " + err.Error())
		}
	}()
	response := bufio.NewWriter(conn)

	welcomeMessage := "SMTP Server for testd9t is ready"
	n, err := fmt.Fprintln(response, "220"+welcomeMessage)
	if err != nil {
		log.Error("Error sending response: " + err.Error())
	} else {
		log.Info("[SMTP] Sent: 220 SMTP Server for testd9t is ready: " + strconv.Itoa(n))
	}

	err = response.Flush()
	if err != nil {
		log.Error("Error flushing response: " + err.Error())
	} else {
		log.Info("[SMTP] Flushed response")
	}

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
			log.Info("[SMTP] Received: 250-Hello")
			response.Flush()
		} else if strings.HasPrefix(line, "AUTH PLAIN") {
			fmt.Fprintln(response, "334 ")
			log.Info("[SMTP] Received: AUTH PLAIN")
			response.Flush()

			if !scanner.Scan() {
				fmt.Fprintln(response, "535 Authentication failed")
				response.Flush()
				break
			}
			decoded, _ := base64.StdEncoding.DecodeString(scanner.Text())
			parts := strings.SplitN(string(decoded), "\x00", 3)
			log.Info("[SMTP] Received: " + string(decoded))
			log.Info("       Received: " + strings.Join(parts, ", "))
			if len(parts) < 3 {
				fmt.Fprintln(response, "535 Need username and password, base64 encoded")
				response.Flush()
				continue
			}

			email := parts[1]
			password := parts[2]

			log.Info("[SMTP] Authenticating with\nemail: " + email + "\npassword" + password)
			_, status := s.repository.GetEmailAddressByName(email)
			if status.Err() {
				log.Error("Could not find email address: " + email)
				fmt.Fprintln(response, "535 Authentication failed")
				response.Flush()
				continue
			}

			log.Info("[SMTP] Comparing if " + password + " = " + os.Getenv(constants.MASTER_PASSWORD))
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
	} else {
		log.Info("[SMTP] Connection closed")
	}
}
