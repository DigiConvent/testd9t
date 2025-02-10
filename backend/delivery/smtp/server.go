package smtp_server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type SMTPServer struct {
	Address string
}

func (s *SMTPServer) HandleConnection(conn net.Conn) {
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
