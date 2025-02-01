package log

import (
	"fmt"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"

// var blue = "\033[34m"
// var magenta = "\033[35m"
var cyan = "\033[36m"

// var gray = "\033[37m"
// var white = "\033[97m"

type Logger struct {
	Level int    `json:"level"`
	Msg   string `json:"msg"`
}

// levels:
// 0 - danger
// 1 - warning
// 2 - info
// 3 - success

var logger = &Logger{}

func SetLogLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > 3 {
		level = 3
	}
	logger.Level = level
}

func Error(msg string) {
	fmt.Println(red + msg + reset)
}

func Warning(msg string) {
	if logger.Level >= 1 {
		fmt.Println(yellow + msg + reset)
	}
}

func Info(msg string) {
	if logger.Level >= 2 {
		fmt.Println(cyan + msg + reset)
	}
}

func Success(msg string) {
	if logger.Level >= 3 {
		fmt.Println(green + msg + reset)
	}
}
