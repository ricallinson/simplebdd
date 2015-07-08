package simplebdd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var TEST_DIR = "test-results"
var FILE_MASK os.FileMode = 0777

type testResult struct {
	name      string
	classname string
	time      float32
}

func writeXunitFile(name string, time float32, tests []testResult, errors int, skipped int, failures int) {
	xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	xml += fmt.Sprintf("<testsuite name=\"%s\" time=\"%f\" tests=\"%d\" errors=\"%d\" skipped=\"%d\" failures=\"%d\">\n", name, time, len(tests), errors, skipped, failures)
	for _, test := range tests {
		xml += fmt.Sprintf("\t<testcase name=\"%s\" classname=\"%s\" time=\"%f\"/>\n", test.name, test.classname, test.time)
	}
	xml += "</testsuite>\n"
	ioutil.WriteFile("./"+TEST_DIR+"/xunit_"+strings.Replace(name, " ", "_", -1)+".xml", []byte(xml), FILE_MASK)
}

func Xunit(t *testing.T) {

	group := ""
	tests := []testResult{}
	passed := 0
	failed := 0
	skipped := 0

	os.Mkdir("./"+TEST_DIR, FILE_MASK)

	for _, test := range testRun.tests {
		if group != test.group {
			if group != "" {
				writeXunitFile(group, 0.0, tests, 0, skipped, failed)
			}
			group = test.group
			tests = []testResult{}
			passed = 0
			failed = 0
			skipped = 0
		}
		if test.skipped {
			skipped++
		} else if test.passed {
			passed++
		} else {
			failed++
		}
		tests = append(tests, testResult{test.group, test.message, 0.0})
	}
	writeXunitFile(group, 0.0, tests, 0, skipped, failed)
}
