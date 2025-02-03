package log

import (
	"fmt"
	"time"

	core_utils "github.com/DigiConvent/testd9t/core/utils"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"

// var blue = "\033[34m"
// var magenta = "\033[35m"
var cyan = "\033[36m"

// var gray = "\033[37m"
var white = "\033[97m"

type Logger struct {
	Level int    `json:"level"`
	Msg   string `json:"msg"`
}

// levels:
// 0 - danger
// 1 - warning
// 2 - info
// 3 - success

func Error(msg string) {
	fmt.Println(white + getTime() + ": " + red + msg + reset)
}

func Warning(msg string) {
	fmt.Println(white + getTime() + ": " + yellow + msg + reset)
}

func Info(msg string) {
	fmt.Println(white + getTime() + ": " + cyan + msg + reset)
}

func Success(msg string) {
	fmt.Println(white + getTime() + ": " + green + msg + reset)
}

func getTime() string {
	return time.Now().Format(core_utils.FormattedTime)
}
