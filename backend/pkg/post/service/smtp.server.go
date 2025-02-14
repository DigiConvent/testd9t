// exempt from testing
package post_service

import (
	"crypto/tls"
	"encoding/base64"
	"net"
	"os"
	"os/signal"
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
		go s.handleConnection(connection, &tlsConfig)
	}
}

func (s *PostService) handleConnection(conn net.Conn, tlsConfig *tls.Config) {
	defer conn.Close()

	sendResponse(conn, "220 Welcome to my SMTP server")
	tlsEnabled := false

	buf := make([]byte, 1024)
	var commandQueue []string

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Warning("Error reading from connection: " + err.Error())
			return
		}

		commands := strings.Split(string(buf[:n]), "\r\n")
		for _, cmd := range commands {
			if cmd == "" {
				continue
			}
			commandQueue = append(commandQueue, cmd)
		}

		for len(commandQueue) > 0 {
			cmd := commandQueue[0]
			commandQueue = commandQueue[1:]

			switch {
			case strings.HasPrefix(cmd, "EHLO"):
				sendResponse(conn, "250-HELLO")
				sendResponse(conn, "250-PIPELINING")
				sendResponse(conn, "250-AUTH PLAIN")
				sendResponse(conn, "250 STARTTLS")
			case strings.HasPrefix(cmd, "AUTH PLAIN"):
				if !tlsEnabled {
					sendResponse(conn, "538 Encryption required for AUTH PLAIN")
					continue
				}

				input := strings.TrimPrefix(cmd, "AUTH PLAIN ")
				decoded, err := base64.StdEncoding.DecodeString(input)
				if err != nil {
					log.Warning("Error decoding AUTH PLAIN input: " + err.Error())
					sendResponse(conn, "535 Could not decode AUTH PLAIN "+input)
					continue
				}

				segments := strings.Split(string(decoded), "\x00")
				if len(segments) != 3 {
					log.Warning("Error parsing AUTH PLAIN input: " + strings.Join(segments, ", "))
					sendResponse(conn, "535 Could not parse AUTH PLAIN "+input)
					continue
				}

				_, status := s.repository.GetEmailAddressByName(strings.Split(segments[1], "@")[0])
				if status.Err() {
					log.Warning("Error getting email address by name: " + status.Message)
					sendResponse(conn, "535 Could not parse AUTH PLAIN "+input)
					continue
				}

				if segments[2] != os.Getenv(constants.MASTER_PASSWORD) {
					log.Warning("Error checking password: got " + segments[2] + ", expected " + os.Getenv(constants.MASTER_PASSWORD))
					sendResponse(conn, "535 Could not parse AUTH PLAIN "+input)
					continue
				}

				sendResponse(conn, "235 OK")
			case strings.HasPrefix(cmd, "STARTTLS"):
				if tlsEnabled {
					sendResponse(conn, "503 TLS already enabled")
					continue
				}
				sendResponse(conn, "220 Ready to start TLS")

				tlsConn := tls.Server(conn, tlsConfig)
				err := tlsConn.Handshake()
				if err != nil {
					log.Warning("TLS handshake failed: " + err.Error())
					return
				}

				conn = tlsConn
				tlsEnabled = true
				log.Info("TLS enabled")
			case strings.HasPrefix(cmd, "MAIL FROM"):
				sendResponse(conn, "250 OK")
			case strings.HasPrefix(cmd, "RCPT TO"):
				sendResponse(conn, "250 OK")
			case strings.HasPrefix(cmd, "DATA"):
				sendResponse(conn, "354 Start mail input; end with <CRLF>.<CRLF>")
				handleData(conn)
			case strings.HasPrefix(cmd, "QUIT"):
				sendResponse(conn, "221 Bye")
				return
			default:
				sendResponse(conn, "502 Command not implemented")
			}
		}
	}
}

func sendResponse(conn net.Conn, response string) {
	_, err := conn.Write([]byte(response + "\r\n"))
	if err != nil {
		log.Warning("Error sending response: " + err.Error())
	}
}

func handleData(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Warning("Error reading data: " + err.Error())
			return
		}

		data := string(buf[:n])
		if strings.Contains(data, "\r\n.\r\n") {
			sendResponse(conn, "250 OK")
			return
		}
	}
}
