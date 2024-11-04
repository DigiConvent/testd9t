package installation_utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteVerbose(input string, returnOutput bool) string {
	output := Execute(input, true)
	fmt.Println(output)
	if returnOutput {
		return output
	}
	return ""
}

func Execute(input string, returnOutput bool) string {
	var output []byte
	cmd := exec.Command(strings.Split(input, " ")[0], strings.Split(input, " ")[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	if returnOutput {
		return string(output)
	}
	return ""
}
