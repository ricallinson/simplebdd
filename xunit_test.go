package simplebdd

import (
    "testing"
    "io/ioutil"
    "strings"
)

func TestXunit(t *testing.T) {

    testsuite := "xunit-test-file"
    testcase := "should return true after being called"

    Describe(testsuite, func() {
        It(testcase, func() {
            AssertEqual(true, true)
        })
    })

    Xunit(t)

    Describe("Xunit() Read", func() {

        xml, _ := ioutil.ReadFile("./" + TEST_DIR + "/xunit-test-file.xml")

        It("should return true because it contains the text " + testsuite, func() {
            AssertEqual(strings.Contains(string(xml), testsuite), true)
        })

        It("should return true because it contains the text " + testcase, func() {
            AssertEqual(strings.Contains(string(xml), testcase), true)
        })
    })

    Report(t)
}
