package mime

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"strings"
)

type EmailContent struct {
	PlainText   string
	HTMLText    string
	Attachments map[string][]byte
}

func ParseEmail(rawEmail string) (*EmailContent, error) {
	msg, err := mail.ReadMessage(strings.NewReader(rawEmail))
	if err != nil {
		log.Fatal(err)
	}

	contentType := msg.Header.Get("Content-Type")
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		log.Fatal("Email is not multipart")
	}

	bodyReader := multipart.NewReader(msg.Body, params["boundary"])

	var plaintext, htmltext string
	attachments := make(map[string][]byte)

	for {
		part, err := bodyReader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		partContentType := part.Header.Get("Content-Type")
		partMediaType, partParams, err := mime.ParseMediaType(partContentType)
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(partMediaType, "multipart/") {
			nestedReader := multipart.NewReader(part, partParams["boundary"])
			for {
				nestedPart, err := nestedReader.NextPart()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				nestedPartContentType := nestedPart.Header.Get("Content-Type")
				nestedPartMediaType, _, err := mime.ParseMediaType(nestedPartContentType)
				if err != nil {
					log.Fatal(err)
				}

				body, err := io.ReadAll(nestedPart)
				if err != nil {
					log.Fatal(err)
				}

				if nestedPartMediaType == "text/plain" {
					plaintext = string(body)
				} else if nestedPartMediaType == "text/html" {
					htmltext = string(body)
				}
			}
		} else {
			body, err := io.ReadAll(part)
			if err != nil {
				log.Fatal(err)
			}

			var decodedBody []byte
			encoding := strings.ToLower(part.Header.Get("Content-Transfer-Encoding"))
			if encoding == "base64" {
				decodedBody, err = base64.StdEncoding.DecodeString(string(body))
				if err != nil {
					log.Fatal(err)
				}
			} else if encoding == "quoted-printable" {
				decodedBody, err = io.ReadAll(quotedprintable.NewReader(bytes.NewReader(body)))
				if err != nil {
					log.Fatal(err)
				}
			} else {
				decodedBody = body
			}

			if partMediaType == "text/plain" {
				plaintext = string(body)
			} else if partMediaType == "text/html" {
				htmltext = string(body)
			}

			contentDisposition := part.Header.Get("Content-Disposition")
			if strings.HasPrefix(contentDisposition, "attachment") {
				_, params, err := mime.ParseMediaType(contentDisposition)
				if err != nil {
					log.Fatal(err)
				}
				filename := params["filename"]
				attachments[filename] = decodedBody
			}
		}
	}

	return &EmailContent{
		PlainText:   plaintext,
		HTMLText:    strings.TrimSpace(htmltext),
		Attachments: attachments,
	}, nil
}
