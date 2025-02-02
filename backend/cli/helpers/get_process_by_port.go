package cli_helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetProcessName(pid int) (string, error) {
	statusFile := fmt.Sprintf("/proc/%d/status", pid)
	file, err := os.Open(statusFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Name:") {
			parts := strings.Fields(scanner.Text())
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}
	return "", fmt.Errorf("could not determine process name")
}

func GetPIDByPort(port int) (int, error) {
	file, err := os.Open("/proc/net/tcp")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 10 {
			continue
		}

		localAddr := fields[1]
		parts := strings.Split(localAddr, ":")
		if len(parts) != 2 {
			continue
		}

		hexPort := parts[1]
		p, err := strconv.ParseInt(hexPort, 16, 32)
		if err != nil {
			continue
		}

		if int(p) == port {
			pidField := fields[9]
			pid, err := strconv.Atoi(pidField)
			if err != nil {
				return 0, err
			}
			return pid, nil
		}
	}
	return 0, fmt.Errorf("no process found using port %d", port)
}
