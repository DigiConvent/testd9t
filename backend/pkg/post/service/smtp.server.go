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
		} else {
			log.Info("[SMTP] Connection from " + conn.RemoteAddr().String() + " closed")
		}
	}()

	welcomeMessage := "SMTP Server for testd9t is ready"
	n, err := conn.Write([]byte("220 " + welcomeMessage + "\r\n"))
	if err != nil {
		log.Info(err.Error())
	}
	if err != nil {
		log.Error("Error sending response: " + err.Error())
	} else {
		log.Info("[SMTP] Sent: 220 SMTP Server for testd9t is ready: " + strconv.Itoa(n))
	}

	scanner := bufio.NewScanner(conn)

	authenticated := false
	isData := false
	var sender, recipient, data string

	for scanner.Scan() {
		line := scanner.Text()
		log.Info("[SMTP] C: " + line)

		if strings.HasPrefix(line, "HELO") || strings.HasPrefix(line, "EHLO") {
			_, err = conn.Write([]byte("250-testd9t-smtp"))
			if err != nil {
				log.Warning("[SMTP] S:" + "250-testd9t-smtp:" + err.Error())
			} else {
				log.Success("[SMTP] S: 250-testd9t-smtp")
			}

			_, err = conn.Write([]byte("250 AUTH PLAIN"))
			if err != nil {
				log.Warning("[SMTP] S: " + "250 AUTH PLAIN LOGIN:" + err.Error())
			} else {
				log.Success("[SMTP] S: 250 AUTH PLAIN")
			}
		} else if strings.HasPrefix(line, "AUTH PLAIN") {
			decoded, _ := base64.StdEncoding.DecodeString(scanner.Text())
			parts := strings.SplitN(string(decoded), "\x00", 3)
			log.Info("[SMTP] Received: " + string(decoded))
			log.Info("       Received: " + strings.Join(parts, ", "))
			if len(parts) < 3 {
				_, err = conn.Write([]byte("535 Need username and password, base64 encoded"))
				if err != nil {
					log.Info("Failed to send " + "535 Need username and password, base64 encoded:" + err.Error())
				}
				continue
			}

			email := parts[1]
			password := parts[2]

			log.Info("[SMTP] Authenticating with\nemail: " + email + "\npassword" + password)
			_, status := s.repository.GetEmailAddressByName(email)
			if status.Err() {
				log.Error("Could not find email address: " + email)
				_, err = conn.Write([]byte("535 Authentication failed"))
				if err != nil {
					log.Info("Failed to send " + "535 Authentication failed:" + err.Error())
				}
				continue
			}

			log.Info("[SMTP] Comparing if " + password + " = " + os.Getenv(constants.MASTER_PASSWORD))
			if password == os.Getenv(constants.MASTER_PASSWORD) {
				authenticated = true
				_, err = conn.Write([]byte("235 Authentication successful"))
				if err != nil {
					log.Info("Failed to send " + "235 Authentication successful:" + err.Error())
				}
			} else {
				_, err = conn.Write([]byte("535 Authentication failed"))
				if err != nil {
					log.Info("Failed to send " + "535 Authentication failed:" + err.Error())
				}
			}
		} else if strings.HasPrefix(line, "MAIL FROM:") {
			if !authenticated {
				_, err = conn.Write([]byte("530 Authentication required"))
				if err != nil {
					log.Info("Failed to send " + "530 Authentication required:" + err.Error())
				}
				continue
			}
			sender = strings.TrimSpace(strings.TrimPrefix(line, "MAIL FROM:"))
			_, err = conn.Write([]byte("250 OK"))
			if err != nil {
				log.Info("Failed to send " + "250 OK:" + err.Error())
			}
		} else if strings.HasPrefix(line, "RCPT TO:") {
			if !authenticated {
				_, err = conn.Write([]byte("530 Authentication required"))
				if err != nil {
					log.Info("Failed to send " + "530 Authentication required:" + err.Error())
				}
				continue
			}
			recipient = strings.TrimSpace(strings.TrimPrefix(line, "RCPT TO:"))
			_, err = conn.Write([]byte("250 OK"))
			if err != nil {
				log.Info("Failed to send " + "250 OK:" + err.Error())
			}
		} else if strings.HasPrefix(line, "DATA") {
			if !authenticated {
				_, err = conn.Write([]byte("530 Authentication required"))
				if err != nil {
					log.Info("Failed to send " + "530 Authentication required:" + err.Error())
				}
				continue
			}
			_, err = conn.Write([]byte("354 End data with <CR><LF>.<CR><LF>"))
			if err != nil {
				log.Info("Failed to send " + "354 End data with <CR><LF>.<CR><LF:>" + err.Error())
			}
			isData = true
			data = ""
		} else if isData {
			if strings.TrimSpace(line) == "." {
				fmt.Println("Received Email:")
				fmt.Println("From:", sender)
				fmt.Println("To:", recipient)
				fmt.Println("Data:", data)
				_, err = conn.Write([]byte("250 OK: Message accepted"))
				if err != nil {
					log.Info("Failed to send " + "250 OK: Message accepted:" + err.Error())
				}
				isData = false
			} else {
				data += line + "\n"
			}
		} else if strings.HasPrefix(line, "QUIT") {
			_, err = conn.Write([]byte("221 Bye"))
			if err != nil {
				log.Info("Failed to send " + "221 Bye:" + err.Error())
			}
			break
		} else {
			_, err = conn.Write([]byte("500 Unrecognized command"))
			if err != nil {
				log.Info("Failed to send " + "500 Unrecognized command:" + err.Error())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("Connection error: " + err.Error())
	} else {
		log.Info("[SMTP] Connection closed")
	}
}
