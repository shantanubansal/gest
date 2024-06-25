package wrapper

import (
	"bufio"
	"log"
	"os"
	"sync"
	"testing"
)

type TestWrapper struct {
	inputTests   map[string]func(t *testing.T)
	allowedTests map[string]bool
	allAllowed   bool
	sw           *sync.Mutex
}

var tw = NewTestWrapper()

func GetTestWrapper() *TestWrapper {
	return tw
}

func AddTest(name string, f func(t *testing.T)) {
	tw.sw.Lock()
	defer tw.sw.Unlock()
	tw.inputTests[name] = f
}
func NewTestWrapper() *TestWrapper {
	tw := &TestWrapper{inputTests: make(map[string]func(t *testing.T)), sw: &sync.Mutex{}}
	//aT, err := LoadAllowedTests()
	//if err != nil {
	//	log.Printf("Unable to load the allowed tests: %v", err)
	//	tw.allAllowed = true
	//}
	tw.allAllowed = true //	tw.allowedTests = aT
	return tw
}

func (tw *TestWrapper) RunAll(t *testing.T) {
	for name, f := range tw.inputTests {
		if tw.allAllowed || tw.allowedTests[name] {
			t.Run(name, f)
		} else {
			log.Printf("Skipping test: %s", name)
		}
	}
}

func LoadAllowedTests() (map[string]bool, error) {
	testTag := os.Getenv("TEST_TAG")
	if testTag == "" {
		return nil, nil
	}

	fileName := "tests_" + testTag + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	allowedTests := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allowedTests[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allowedTests, nil
}
