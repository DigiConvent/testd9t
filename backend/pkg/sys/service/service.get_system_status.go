// exempt from testing
package sys_service

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"blitiri.com.ar/go/spf"
	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/go-telegram/bot"
)

func (s *SysService) GetSystemStatus() (*sys_domain.SystemStatus, *core.Status) {
	systemStatus := &sys_domain.SystemStatus{}

	programVersion, databaseVersion, status := s.repository.GetCurrentVersion()
	if status.Err() {
		systemStatus.ProgramVersion = sys_domain.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
		systemStatus.DatabaseVersion = sys_domain.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
	} else {
		systemStatus.ProgramVersion = *programVersion
		systemStatus.DatabaseVersion = *databaseVersion
	}

	var stat syscall.Statfs_t
	var testd9tStat syscall.Statfs_t

	syscall.Statfs("/", &stat)
	syscall.Statfs(constants.HOME_PATH, &testd9tStat)

	systemStatus.TotalSpace = stat.Blocks * uint64(stat.Bsize)
	systemStatus.FreeSpace = stat.Bfree * uint64(stat.Bsize)

	cmd := exec.Command("du", "-sb", constants.HOME_PATH)
	out, _ := cmd.Output()
	outString := string(out)
	outBytes, _ := strconv.Atoi(outString)
	systemStatus.DataSpace = uint64(outBytes)

	botToken, status := s.repository.GetBotToken()
	if status.Err() || botToken == "" {
		systemStatus.TelegramBotStatus = "false"
		systemStatus.TelegramBotHint = "No Telegram Bot Token not found"
	} else {
		bot, err := bot.New(botToken)
		if err != nil {
			systemStatus.TelegramBotStatus = "false"
			systemStatus.TelegramBotHint = err.Error()
		} else {
			defer bot.Close(context.TODO())
			systemStatus.TelegramBotStatus = botToken[0:4] + "...." + botToken[len(botToken)-4:]
			systemStatus.TelegramBotHint = "Telegram bot token found and working"
		}
	}

	ips, err := net.LookupHost(os.Getenv(constants.DOMAIN))
	if err != nil || len(ips) == 0 {
		systemStatus.DnsStatus = "false"
		systemStatus.DnsHint = "Add an A record for " + os.Getenv(constants.DOMAIN)
	} else {
		systemStatus.DnsStatus = strings.Join(ips, ", ")
		mxStatus, err := net.LookupMX(os.Getenv(constants.DOMAIN))
		if err != nil {
			systemStatus.MxStatus = "false"
			systemStatus.MxHint = err.Error()
		} else {
			if len(mxStatus) == 0 {
				systemStatus.MxStatus = "false"
				systemStatus.MxHint = "MX record not found"
			} else {
				systemStatus.MxStatus = mxStatus[0].Host
				systemStatus.MxHint = "MX record present"
			}
		}
	}

	txtRecords, err := net.LookupTXT(os.Getenv(constants.DOMAIN))
	if err != nil {
		systemStatus.DkimStatus = "false"
		systemStatus.DkimHint = err.Error()
		systemStatus.SpfStatus = "false"
		systemStatus.SpfHint = err.Error()
	} else {
		if len(txtRecords) == 0 {
			systemStatus.DkimStatus = "false"
			systemStatus.DkimHint = "DKIM record not found"
		} else {
			for _, value := range txtRecords {
				// https://datatracker.ietf.org/doc/html/rfc6376/#section-7.5
				if strings.HasPrefix(value, "v=DKIM1") {
					systemStatus.DkimStatus, systemStatus.DkimHint = getDkimStatus(value)
				}
				// https://datatracker.ietf.org/doc/html/rfc7208#section-4.5
				if strings.HasPrefix(value, "v=spf1") {
					systemStatus.SpfStatus, systemStatus.SpfHint = getSpfStatus(value, ips)
				}
			}
		}
	}

	return systemStatus, core.StatusSuccess()
}

// https://datatracker.ietf.org/doc/html/rfc7208#section-4.6.1
func getSpfStatus(value string, ips []string) (string, string) {
	caughtErrors := make([]string, 0)

	for _, rawIp := range ips {
		parsedIp := net.ParseIP(rawIp)
		result, err := spf.CheckHostWithSender(parsedIp, "", os.Getenv(constants.DOMAIN))
		if err != nil {
			caughtErrors = append(caughtErrors, err.Error())
		}
		if result == spf.Fail {
			caughtErrors = append(caughtErrors, "Check for "+rawIp+" failed")
		}
	}

	if len(caughtErrors) == 0 {
		return value, "SPF record valid"
	}

	return value, strings.Join(caughtErrors, ", ")
}

func getDkimStatus(value string) (string, string) {
	result := make(map[string]string)
	parts := strings.Split(value, "; ")

	for _, part := range parts {
		kv := strings.Split(part, "=")
		key := kv[0]
		val := kv[1]

		result[key] = val
	}

	rawPublicKey, err := os.ReadFile(post_setup.DkimPublicKeyPath())

	if err != nil {
		return "false", "Dkim is not set up properly on the server"
	}

	block, _ := pem.Decode(rawPublicKey)
	if block == nil {
		return "false", "Could not decode dkim key"
	}

	parsedPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "false", "Could not parse public key"
	}

	derFormat, err := x509.MarshalPKIXPublicKey(parsedPublicKey)

	if err != nil {
		return "false", "Could not marshal public key"
	}

	base64PublicKey := base64.StdEncoding.EncodeToString(derFormat)
	dkimShouldBe := "Add the following TXT record: _default._domainkey." + os.Getenv(constants.DOMAIN) + " with value 'v=DKIM1; k=rsa; p=" + base64PublicKey + ";'"

	if result["p"] != base64PublicKey || result["k"] != "rsa" {
		return "false", dkimShouldBe
	}

	return dkimShouldBe, "valid"
}
