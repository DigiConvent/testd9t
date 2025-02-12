package log

import (
	"fmt"
	"strings"
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

func Error(msg interface{}) {
	fmt.Println(white+getTime()+": "+red, prep(msg), reset)
}

func Warning(msg interface{}) {
	fmt.Println(white+getTime()+": "+yellow, prep(msg), reset)
}

func Info(msg interface{}) {
	fmt.Println(white+getTime()+": "+cyan, prep(msg), reset)
}

func Success(msg interface{}) {
	fmt.Println(white+getTime()+": "+green, prep(msg), reset)
}

func getTime() string {
	return time.Now().Format(core_utils.FormattedTime)
}

func prep(input interface{}) string {
	stringInput := fmt.Sprint(input)
	segments := strings.Split(stringInput, "\n")
	for i := range segments {
		if segments[i] == "" {
			continue
		}
		segments[i] = strings.Repeat(" ", 19) + segments[i]
	}
	return strings.Join(segments, "\n")
}
