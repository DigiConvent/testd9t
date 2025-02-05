package main_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

type TestEvent struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
}

type TestPackage struct {
	Name   string `json:"Name"`
	events []TestEvent
}

func (p TestPackage) GetFailingTests() []TestEvent {
	var failingTests []TestEvent
	for _, event := range p.events {
		if event.Action == "fail" {
			failingTests = append(failingTests, event)
		}
	}
	return failingTests
}

func (p TestPackage) GetPassingTests() []TestEvent {
	var passingTests []TestEvent
	for _, event := range p.events {
		if event.Action == "pass" {
			passingTests = append(passingTests, event)
		}
	}
	return passingTests
}

func (p TestPackage) GetSummary(spaces int) string {
	if len(p.events) == 0 {
		return ""
	}
	var result string
	failingTests := p.GetFailingTests()

	var success = fmt.Sprintf("%s%3d%s", green, len(p.events)-len(failingTests), reset)

	var failing string
	if len(failingTests) != 0 {
		failing = fmt.Sprintf("%s%s%3d%s", red, bold, len(failingTests), reset)
	} else {
		failing = fmt.Sprintf("%s%3d%s", red, len(failingTests), reset)
	}
	result += fmt.Sprintf("%s%"+strconv.Itoa(spaces)+"s%s %4s %4s %s\n", bold, p.Name, reset, success, failing, reset)

	if len(failingTests) != 0 {
		result += "Fix the following cases\n"
	}
	for _, event := range failingTests {
		result += fmt.Sprintf("%s%s%s\n", red, p.Name+":"+event.Test, reset)
	}

	result = strings.TrimSuffix(result, "\n")

	return result
}

const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
	bold  = "\033[1m"
)

func TestMain(m *testing.M) {
	main()
}

func main() {
	if os.Getenv("TEST_RUNNER") == "1" {
		return
	}

	os.Setenv("TEST_RUNNER", "1")

	cmd := exec.Command("go", "test", "./...", "-json", "-v")
	cmd.Env = append(os.Environ(), "TEST_RUNNER=1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	packages := make(map[string]TestPackage)

	scanner := bufio.NewScanner(stdout)

	toIgnore := map[string]bool{"start": true, "run": true, "pass": true, "skip": true}

	for scanner.Scan() {
		var event TestEvent
		line := scanner.Text()
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			continue
		}

		if _, ok := packages[event.Package]; !ok {
			packages[event.Package] = TestPackage{Name: event.Package}
		}

		if _, ok := toIgnore[event.Action]; ok {
			continue
		}

		if strings.Contains(event.Output, "no test files") {
			continue
		}

		if event.Test == "" {
			continue
		}

		events := packages[event.Package].events
		events = append(events, event)
		packages[event.Package] = TestPackage{Name: event.Package, events: events}
	}

	_ = cmd.Wait()

	maxPkgNameLen := 0
	hasErrors := false

	for _, pkg := range packages {
		if len(pkg.events) == 0 {
			continue
		}
		if len(pkg.GetFailingTests()) != 0 {
			hasErrors = true
		}
		if len(pkg.Name) > maxPkgNameLen {
			maxPkgNameLen = len(pkg.Name)
		}
	}

	for _, pkg := range packages {
		summary := pkg.GetSummary(maxPkgNameLen)
		if summary == "" {
			continue
		}
		fmt.Println(summary)
	}

	if hasErrors {
		os.Exit(1)
	}
	os.Exit(0)
}
