package main_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"testing"
)

const (
	green  = "\033[32m"
	red    = "\033[31m"
	reset  = "\033[0m"
	bold   = "\033[1m"
	orange = "\033[93m"
)

func TestMain(m *testing.M) {
	for _, arg := range os.Args {
		if strings.HasSuffix(arg, "testlog.txt") {
			return
		}
	}

	fmt.Println("Running tests")
	result := true
	result = main() && result
	for _, missing := range CheckForMissingTests() {
		fmt.Println("Missing test file for", missing)
		result = false
	}
	if !result {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

type TestEvent struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
}

func (e *TestEvent) filename() string {
	pkgName, _ := strings.CutPrefix(e.Package, "github.com/DigiConvent/testd9t/")
	return FindTestFile(pkgName, e.Test)
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

func (p *TestPackage) GetSkippedTests() []TestEvent {
	var skippedTests []TestEvent
	for _, event := range p.events {
		if strings.Contains(event.Output, "--- SKIP") {
			skippedTests = append(skippedTests, event)
		}
	}
	return skippedTests
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

	var success = fmt.Sprintf("%s%8d%s", green, len(p.GetPassingTests()), reset)
	var skipped = fmt.Sprintf("%s%8d%s", orange, len(p.GetSkippedTests()), reset)

	var failing string
	if len(failingTests) != 0 {
		failing = fmt.Sprintf("%s%s%8d%s", red, bold, len(failingTests), reset)
	} else {
		failing = fmt.Sprintf("%s%8d%s", red, len(failingTests), reset)
	}
	result += fmt.Sprintf("%-"+strconv.Itoa(spaces)+"s%s %s %s %s %s\n", p.Name, reset, success, skipped, failing, reset)

	result = strings.TrimSuffix(result, "\n")

	return result
}

func main() bool {
	if os.Getenv("TEST_RUNNER") == "1" {
		return false
	}

	os.Setenv("TEST_RUNNER", "1")

	cmd := exec.Command("go", "test", "./...", "-json", "-v")
	cmd.Env = append(os.Environ(), "TEST_RUNNER=1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
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

	totalFailing := 0
	totalPassing := 0
	totalSkipped := 0
	allFailingTests := make([]TestEvent, 0)
	fmt.Printf("%-"+strconv.Itoa(maxPkgNameLen)+"s %8s %8s %8s\n", "Package", "Passed", "Skipped", "Failed")
	for pkgName := range sortedByLength {
		pkg := packages[sortedByLength[pkgName]]
		if len(pkg.events) == 0 {
			continue
		}
		failingTests := pkg.GetFailingTests()
		allFailingTests = append(allFailingTests, failingTests...)
		totalFailing += len(failingTests)
		totalPassing += len(pkg.GetPassingTests())
		totalSkipped += len(pkg.GetSkippedTests())
		summary := pkg.GetSummary(maxPkgNameLen)
		if summary == "" {
			continue
		}
		fmt.Println(summary)
	}

	fmt.Println(strings.Repeat("-", maxPkgNameLen+20))
	fmt.Printf("%-"+strconv.Itoa(maxPkgNameLen)+"s %s%8d%s %s%8d%s %s%8d%s\n", "Total", green, totalPassing, reset, orange, totalSkipped, reset, red, totalFailing, reset)

	for i := range allFailingTests {
		for j := i + 1; j < len(allFailingTests); j++ {
			if len(allFailingTests[i].filename()) > len(allFailingTests[j].filename()) {
				allFailingTests[i], allFailingTests[j] = allFailingTests[j], allFailingTests[i]
			}
		}
	}

	for _, event := range allFailingTests {
		cleanOutput := ""
		segments := strings.Split(event.Output, "\n")
		for _, segment := range segments {
			if strings.Contains(segment, "--- FAIL") || strings.Contains(segment, "=== RUN") {
				continue
			}
			if segment == "" {
				continue
			}
			cleanOutput += fmt.Sprintf("%s%s%s\n", red, segment, reset)
		}

		fmt.Printf("%s%s%s%s\n", bold, red, event.filename(), reset)
		if cleanOutput != "" {
			fmt.Printf("%s'%s'%s\n", red, cleanOutput, reset)
		}
	}

	err = os.WriteFile("test.json", []byte(strconv.Itoa(count)), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if totalFailing > 0 {
		return false
	}
	return true
}

func CheckForMissingTests() []string {
	dirPath, _ := os.Getwd()
	packageFolder, _ := os.ReadDir(path.Join(dirPath, "pkg"))

	missingTestFiles := []string{}
	for _, pkgName := range packageFolder {
		servicesFolder := path.Join(dirPath, "pkg", pkgName.Name(), "service")
		serviceFiles, _ := os.ReadDir(servicesFolder)
		for _, serviceFile := range serviceFiles {
			if serviceFile.IsDir() {
				continue
			}
			if strings.HasSuffix(serviceFile.Name(), "_test.go") {
				continue
			}
			contents, _ := os.ReadFile(path.Join(servicesFolder, serviceFile.Name()))
			if strings.HasPrefix(string(contents), "// exempt from testing") {
				continue
			}

			filename := strings.TrimSuffix(serviceFile.Name(), ".go")
			if _, err := os.Stat(path.Join(servicesFolder, filename+"_test.go")); os.IsNotExist(err) {
				var functionName string
				for _, line := range strings.Split(string(contents), "\n") {
					if strings.Contains(line, "func") && strings.Contains(line, "*core.Status") {
						functionName = strings.Split(line, "Service)")[1]
						functionName = strings.Split(functionName, "(")[0]
						functionName = strings.TrimSpace(functionName)
						break
					}
				}
				os.WriteFile(path.Join(servicesFolder, filename+"_test.go"), []byte("package "+pkgName.Name()+"_service_test\n\nimport \"testing\"\n\nfunc Test"+functionName+"(t *testing.T) {\n\tt.Fail()\n}\n"), 0644)
				missingTestFiles = append(missingTestFiles, serviceFile.Name())
			}
		}
	}

	return missingTestFiles
}

// I need another helper function to find the filename of the file that contains a test function
func FindTestFile(pkgName, functionName string) string {
	dirPath, _ := os.Getwd()
	pkgDir := path.Join(dirPath, pkgName)
	files, _ := os.ReadDir(pkgDir)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), "_test.go") {
			contents, _ := os.ReadFile(path.Join(pkgDir, file.Name()))
			if strings.Contains(string(contents), functionName) {
				return pkgDir + "/" + file.Name()
			}
		}
	}
	return "unknown file"
}
