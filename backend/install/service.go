package install

import (
	"os"
	"os/exec"
	"strings"
)

func StopService() error {
	cmd := exec.Command("systemctl", "stop", "digiconvent")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CreateService() {
	service := `[Unit]
Description=digiconvent Service
After=network.target

[Service]
ExecStart=/usr/local/bin/digiconvent
WorkingDirectory=/home/digiconvent/
Restart=always
User=digiconvent
Group=digiconvent_group
EnvironmentFile=/etc/digiconvent/env
StandardOutput=journal
StandardError=journal
AmbientCapabilities=CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
`
	serviceFile := "/etc/systemd/system/digiconvent.service"
	f, err := os.Create(serviceFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	service = strings.ReplaceAll(service, "$APP_NAME", "digiconvent")
}
