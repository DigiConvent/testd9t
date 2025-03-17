// exempt from testing
package sys_service

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"net"
	"os"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/go-telegram/bot"
)

func (s *SysService) GetSystemStatus() (*sys_domain.SystemStatus, *core.Status) {
	size, _ := s.repository.GetDiskUsage()
	systemStatus := &sys_domain.SystemStatus{
		Space: *size,
	}

	programVersion, databaseVersion, status := s.repository.GetCurrentVersion()
	if status.Err() {
		systemStatus.Version.ProgramVersion = sys_domain.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
		systemStatus.Version.DatabaseVersion = sys_domain.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
	} else {
		systemStatus.Version.ProgramVersion = *programVersion
		systemStatus.Version.DatabaseVersion = *databaseVersion
	}

	botToken, status := s.repository.GetBotToken()
	if status.Err() || botToken == "" {
		systemStatus.TelegramBot.TelegramBotStatus = "false"
		systemStatus.TelegramBot.TelegramBotHint = ""
	} else {
		bot, err := bot.New(botToken)
		if err != nil {
			systemStatus.TelegramBot.TelegramBotStatus = "false"
			systemStatus.TelegramBot.TelegramBotHint = err.Error()
		} else {
			defer bot.Close(context.TODO())
			systemStatus.TelegramBot.TelegramBotStatus = botToken[0:4] + "...." + botToken[len(botToken)-4:]
			systemStatus.TelegramBot.TelegramBotHint = ""
		}
	}

	ipAddress, err := getOutboundIP()
	if err != nil {
		systemStatus.Dns.DnsShould = "Error: Could not find public ip address for this server"
	} else {
		systemStatus.Dns.DnsShould = "A///" + os.Getenv(constants.DOMAIN) + "///" + ipAddress
	}

	ips, err := net.LookupHost(os.Getenv(constants.DOMAIN))
	if err != nil || len(ips) == 0 {
		systemStatus.Dns.DnsIs = ""
	} else {
		systemStatus.Dns.MxShould = "MX///" + os.Getenv(constants.DOMAIN) + "///" + os.Getenv(constants.DOMAIN) + "."
		systemStatus.Dns.DnsIs = "A///" + os.Getenv(constants.DOMAIN) + "///" + strings.Join(ips, ", ")
		mxStatus, err := net.LookupMX(os.Getenv(constants.DOMAIN))
		if err != nil {
			systemStatus.Dns.MxIs = ""
		} else {
			if len(mxStatus) == 0 {
				systemStatus.Dns.MxIs = ""
			} else {
				systemStatus.Dns.MxIs = "MX///" + os.Getenv(constants.DOMAIN) + "///" + mxStatus[0].Host
			}
		}
	}

	dkimShould, err := getRecommendedDkimSignature()
	if err != nil {
		dkimShould = "Error generating a DKIM signature: " + err.Error()
	}
	systemStatus.Dns.DkimShould = dkimShould
	systemStatus.Dns.SpfShould = getRecommendedSpfSignature()
	systemStatus.Dns.DmarcShould = "TXT///" + os.Getenv(constants.DOMAIN) + "///v=DMARC1; p=reject; adkim=s; aspf=s;"

	txtRecords, err := net.LookupTXT(os.Getenv(constants.DOMAIN))
	if err != nil {
		systemStatus.Dns.DkimIs = err.Error()
		systemStatus.Dns.SpfIs = err.Error()
	} else {
		if len(txtRecords) != 0 {
			for _, value := range txtRecords {
				// https://datatracker.ietf.org/doc/html/rfc7208#section-4.5
				if strings.HasPrefix(value, "v=spf1") {
					systemStatus.Dns.SpfIs = "TXT///" + os.Getenv(constants.DOMAIN) + "///" + value
				}
				if strings.HasPrefix(value, "v=DMARC1") {
					systemStatus.Dns.DmarcIs = "TXT///" + os.Getenv(constants.DOMAIN) + "///" + value
				}
			}
		}
	}

	dkimRecords, err := net.LookupTXT(constants.DkimPrefix + os.Getenv(constants.DOMAIN))
	if err != nil {
		systemStatus.Dns.DkimIs = err.Error()
	} else {
		if len(dkimRecords) != 0 {
			for _, value := range dkimRecords {
				// https://datatracker.ietf.org/doc/html/rfc6376/#section-7.5
				if strings.HasPrefix(value, "v=DKIM1") {
					systemStatus.Dns.DkimIs = "TXT///" + constants.DkimPrefix + os.Getenv(constants.DOMAIN) + "///" + value
				}
			}
		}
	}

	systemStatus.Version.BuiltAt = sys_domain.CompiledAt
	systemStatus.Version.OnlineSince = sys_domain.StartTime.Format(core_utils.FormattedTime)

	return systemStatus, core.StatusSuccess()
}

var base64PublicKey string = ""

func getBase64PublicKey() (string, error) {
	if base64PublicKey == "" {
		log.Info("test")
		rawPublicKey, err := os.ReadFile(post_setup.DkimPublicKeyPath())

		if err != nil {
			return "", errors.New("dkim is not set up properly on the server: " + err.Error())
		}

		block, _ := pem.Decode(rawPublicKey)
		if block == nil {
			return "", errors.New("could not decode dkim key")
		}

		parsedPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return "", errors.New("could not parse public key: " + err.Error())
		}

		derFormat, err := x509.MarshalPKIXPublicKey(parsedPublicKey)

		if err != nil {
			return "", errors.New("could not marshal public key")
		}
		base64PublicKey = base64.StdEncoding.EncodeToString(derFormat)
	}
	return base64PublicKey, nil
}

func getRecommendedDkimSignature() (string, error) {
	base64PublicKey, err := getBase64PublicKey()
	if err != nil {
		return "", err
	}

	dkimShouldBe := "TXT///" + constants.DkimPrefix + os.Getenv(constants.DOMAIN) + "///v=DKIM1; k=rsa; p=" + base64PublicKey + ";"

	return dkimShouldBe, nil
}

func getRecommendedSpfSignature() string {
	return "TXT///" + os.Getenv(constants.DOMAIN) + "///v=spf1 mx include:" + os.Getenv(constants.DOMAIN) + " -all"
}

func getOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
