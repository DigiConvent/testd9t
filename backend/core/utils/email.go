package core_utils

import (
	"net"
	"strings"
	"unicode"
)

func ValidEmail(email string) bool {
	email = strings.ToLower(email)
	segments := strings.Split(email, "@")
	if len(segments) != 2 {
		return false
	}

	name := segments[0]
	domain := segments[1]
	if name == "" || domain == "" {
		return false
	}

	if len(name) > 64 {
		return false
	}

	if len(domain) > 255 {
		return false
	}

	if name[0] == '.' || name[len(name)-1] == '.' {
		return false
	}

	if strings.Contains(domain, "..") {
		return false
	}

	if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
		if !checkQuotedName(name) {
			return false
		}
	} else if !checkUnquotedName(name) {
		return false
	}

	if strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		ip := domain[1 : len(domain)-1]

		if strings.HasPrefix(ip, "IPv6:") {
			return net.ParseIP(strings.TrimPrefix(ip, "IPv6:")) != nil
		}

		return net.ParseIP(ip) != nil
	}

	segments = strings.Split(domain, ".")
	for _, label := range segments {
		if len(label) == 0 {
			return false
		}
		if strings.HasPrefix(label, "-") || strings.HasSuffix(label, "-") {
			return false
		}
		for _, r := range label {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' {
				return false
			}
		}
	}

	return true
}

func checkUnquotedName(name string) bool {
	// check if the name is valid
	allowedChars := "abcdefghijklmnopqrstuvwxyz0123456789!#$%&'*+-/=?^_`{|}~."
	if strings.Contains(name, "..") {
		return false
	}

	for i := 0; i < len(name); i++ {
		if !strings.ContainsRune(allowedChars, rune(name[i])) {
			return false
		}
	}
	return true
}

func checkQuotedName(name string) bool {
	// If quoted, it may contain Space, Horizontal Tab (HT), any ASCII graphic except Backslash and Quote and a quoted-pair consisting of a Backslash followed by HT, Space or any ASCII graphic; it may also be split between lines anywhere that HT or Space appears. In contrast to unquoted local-parts, the addresses ".John.Doe"@example.com, "John.Doe."@example.com and "John..Doe"@example.com are allowed.
	allowedChars := " 	"
	forbiddenChars := "\\\""

	if strings.ContainsAny(name, forbiddenChars) {
		return false
	}

	for _, c := range name {
		if strings.ContainsRune(allowedChars, c) {
			continue
		}
		if c > unicode.MaxASCII {
			return false
		}
		if c == '\r' || c == '\n' {
			continue
		}

		if c == '\\' {
			if i := strings.IndexRune(name, '\\'); i != len(name)-1 {
				return false
			}
		}
	}

	return true
}
