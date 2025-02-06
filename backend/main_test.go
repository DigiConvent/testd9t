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

func (p *TestPackage) GetFailingTests() []TestEvent {
	var failingTests []TestEvent
	for _, event := range p.events {
		if strings.Contains(event.Output, "--- FAIL") {
			failingTests = append(failingTests, event)
		}
	}
	return failingTests
}

func (p *TestPackage) GetPassingTests() []TestEvent {
	var passingTests []TestEvent
	for _, event := range p.events {
		if strings.Contains(event.Output, "--- PASS") {
			passingTests = append(passingTests, event)
		}
	}
	return passingTests
}

func (p *TestPackage) AddEvent(event TestEvent) bool {
	for i, e := range p.events {
		if e.Test == event.Test {
			p.events[i].Output += event.Output
			return false
		}
	}

	p.events = append(p.events, event)
	return true
}

func (p *TestPackage) GetSummary(spaces int) string {
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
	result += fmt.Sprintf("%-"+strconv.Itoa(spaces)+"s%s %4s %4s %s\n", p.Name, reset, success, failing, reset)

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
	fmt.Println("Running tests")
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

	packages := make(map[string]*TestPackage)

	scanner := bufio.NewScanner(stdout)

	toIgnore := map[string]bool{"start": true, "run": true, "pass": true, "skip": true}

	count := 0
	lastCount := "0"

	testDataFile, err := os.ReadFile("test.json")
	if err != nil {
		lastCount = "?"
	} else {
		lastCount = strings.TrimSpace(string(testDataFile))
	}

	maxPkgNameLen := 0
	for scanner.Scan() {
		var event TestEvent
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &event)
		if err != nil {
			continue
		} else {
		}

		if _, ok := packages[event.Package]; !ok {
			packages[event.Package] = &TestPackage{Name: event.Package}
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

		if len(event.Package) > maxPkgNameLen {
			maxPkgNameLen = len(event.Package)
		}

		fmt.Printf("\rRunning %d/%s %s", count, lastCount, event.Package)
		pkg := packages[event.Package]
		if pkg.AddEvent(event) {
			count++
		}
	}
	fmt.Printf("\rRan %d/%s tests%s\n", count, lastCount, strings.Repeat(" ", maxPkgNameLen))

	_ = cmd.Wait()

	totalFailing := 0
	totalPassing := 0

	sortedByLength := make([]string, 0, len(packages))

	for pkgName := range packages {
		sortedByLength = append(sortedByLength, pkgName)
	}
	for i := 0; i < len(sortedByLength); i++ {
		for j := i + 1; j < len(sortedByLength); j++ {
			if len(packages[sortedByLength[i]].Name) > len(packages[sortedByLength[j]].Name) {
				sortedByLength[i], sortedByLength[j] = sortedByLength[j], sortedByLength[i]
			}
		}
	}

	for i := 0; i < len(sortedByLength); i++ {
		for j := i + 1; j < len(sortedByLength); j++ {
			if len(packages[sortedByLength[i]].Name) == len(packages[sortedByLength[j]].Name) {
				if packages[sortedByLength[i]].Name > packages[sortedByLength[j]].Name {
					sortedByLength[i], sortedByLength[j] = sortedByLength[j], sortedByLength[i]
				}
			}
		}
	}

	allFailingTests := make([]TestEvent, 0)
	for pkgName := range sortedByLength {
		pkg := packages[sortedByLength[pkgName]]
		failingTests := pkg.GetFailingTests()
		allFailingTests = append(allFailingTests, failingTests...)
		totalFailing += len(failingTests)
		totalPassing += len(pkg.GetPassingTests())
		summary := pkg.GetSummary(maxPkgNameLen)
		if summary == "" {
			continue
		}
		fmt.Println(summary)
	}

	fmt.Println(strings.Repeat("-", maxPkgNameLen+20))
	fmt.Printf("%-"+strconv.Itoa(maxPkgNameLen)+"s %s%3d%s %s%3d%s\n", "Total", green, totalPassing, reset, red, totalFailing, reset)

	for _, event := range allFailingTests {
		cleanOutput := ""
		segments := strings.Split(event.Output, "\n")
		for _, segment := range segments {
			if strings.Contains(segment, "--- FAIL") || strings.Contains(segment, "=== RUN") {
				continue
			}
			cleanOutput += fmt.Sprintf("%s%s%s\n", red, segment, reset)
		}
		fmt.Printf("%s%s%s\n", bold, event.Package+":"+event.Test, reset)
		fmt.Printf("%s%s%s\n", red, cleanOutput, reset)
	}

	err = os.WriteFile("test.json", []byte(strconv.Itoa(count)), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if totalFailing > 0 {
		os.Exit(1)
	}
	os.Exit(0)
}
