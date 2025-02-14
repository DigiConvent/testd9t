// exempt from testing
package post_service

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"mime/quotedprintable"
	"net"
	"net/mail"
	"strings"

	"github.com/DigiConvent/testd9t/core/log"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func (s *PostService) smtpReceiveServer() {
	listener, err := net.Listen("tcp", ":25")
	if err != nil {
		log.Error("Error starting server:" + err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("SMTP server listening on port 25...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("Error accepting connection:" + err.Error())
			continue
		}

		go s.handleSMTPConnection(conn)
	}
}

func (s *PostService) handleSMTPConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintf(conn, "220 Welcome to testd9t mailserver\r\n")

	scanner := bufio.NewScanner(conn)

	var from, to, subject string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Received:", line)

		if strings.HasPrefix(line, "HELO") {
			fmt.Fprintf(conn, "250\r\n")
		} else if strings.HasPrefix(line, "MAIL FROM") {
			from = extractEmail(line)
			fmt.Fprintf(conn, "250 OK\r\n")
		} else if strings.HasPrefix(line, "RCPT TO") {
			to = extractEmail(line)
			fmt.Fprintf(conn, "250 OK\r\n")
		} else if line == "DATA" {
			fmt.Fprintf(conn, "354 Start mail input; end with <CRLF>.<CRLF>\r\n")
			data := ""

			for scanner.Scan() {
				line := scanner.Text()
				if line == "." {
					break
				}
				if strings.HasPrefix(line, "Subject:") {
					subject = line[8:]
				}
				data += line + "\n"
			}

			plain, html, attachments, notes := extractEmailContents(data)

			if html != "" {
				plain = html
			}

			status := s.repository.StoreEmail(&post_domain.EmailWrite{
				From:        from,
				To:          to,
				Subject:     subject,
				Body:        plain,
				Attachments: attachments,
				Notes:       notes,
			})

			if status.Err() {
				log.Error("Error storing email: " + status.Message)
				break
			}

			fmt.Fprintf(conn, "250 OK: Message accepted\r\n")
		} else if line == "QUIT" {
			fmt.Fprintf(conn, "221 Bye\r\n")
			break
		} else {
			fmt.Fprintf(conn, "500 Command unrecognized\r\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("Error reading:" + err.Error())
	}
}

func extractEmail(line string) string {
	var email string
	start := strings.Index(line, "<")
	end := strings.Index(line, ">")
	if start == -1 || end == -1 || start >= end {
		email = ""
	} else {
		email = line[start+1 : end]
	}

	return email
}

func extractEmailContents(raw string) (string, string, map[string][]byte, []string) {
	msg, err := mail.ReadMessage(strings.NewReader(raw))
	if err != nil {
		return "", "", nil, nil
	}

	var textContent, htmlContent string
	attachments := make(map[string][]byte)

	var notes []string
	contentType := msg.Header.Get("Content-Type")
	if strings.Contains(contentType, "multipart") {
		mr := multipart.NewReader(msg.Body, msg.Header.Get("Boundary"))
		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			} else {
				log.Info("Part: " + part.Header.Get("Content-Type") + " (" + part.FileName() + ")")
			}

			if part.Header.Get("Content-Type") == "text/plain" {
				data, _ := io.ReadAll(part)
				textContent = string(data)
			} else if part.Header.Get("Content-Type") == "text/html" {
				data, _ := io.ReadAll(part)
				htmlContent = string(data)
			} else if strings.Contains(part.Header.Get("Content-Disposition"), "attachment") {
				filename := part.FileName()
				if filename != "" {
					encoding := part.Header.Get("Content-Transfer-Encoding")

					rawAttachmentData, _ := io.ReadAll(part)
					var attachmentData []byte
					if encoding == "base64" {
						attachmentData, err = base64.StdEncoding.DecodeString(string(rawAttachmentData))
						if err != nil {
							notes = append(notes, fmt.Sprintf("Failed to decode attachment %s: %s", filename, err.Error()))
						}
					} else if encoding == "quoted-printable" {
						attachmentData, err = io.ReadAll(quotedprintable.NewReader(strings.NewReader(string(rawAttachmentData))))
						if err != nil {
							notes = append(notes, fmt.Sprintf("Failed to decode attachment %s: %s", filename, err.Error()))
						}
					} else {
						attachmentData = []byte(string(rawAttachmentData))
					}
					attachments[filename] = attachmentData
				}
			}
		}
	} else {
		data, err := io.ReadAll(msg.Body)
		if err != nil {
			return "", "", nil, notes
		}
		textContent = string(data)
	}

	log.Info("Notes: " + strings.Join(notes, ", "))

	return textContent, htmlContent, attachments, notes
}
