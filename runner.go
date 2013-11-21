/*
    Super simple behavior-driven development style test writer for Go

        package something

        import(
            "testing"
            . "github.com/ricallinson/simplebdd"
        )

        func TestSomething(t *testing.T) {
            Describe("AssertEqual()", func() {
                It("should return that true is true", func() {
                    AssertEqual(true, true)
                })
            })
            Describe("AssertNotEqual()", func() {
                It("should return that false is not true", func() {
                    AssertNotEqual(false, true)
                })
            })
            Report(t)
        }
*/
package simplebdd

import(
    "fmt"
)

/*
    result type
*/

type result struct {
    group string
    message string
    expect string
    result string
    passed bool
}

/*
    testRunner type.
*/

type testRunner struct {
    describe string
    before func()
    beforeEach func()
    it string
    after func()
    afterEach func()
    tests []result
}

/*
    Holds all the tests run.
*/

var testRun = testRunner{}

/*
    Store the result of an Assert call.
*/

func record(pass bool, got interface{}, expected interface{}) {
    testRun.tests = append(testRun.tests, result{
        group: testRun.describe,
        message: testRun.it,
        expect: fmt.Sprint(expected),
        result: fmt.Sprint(got),
        passed: pass,
    })
}

/*
    Called before each Describe() function call.
*/
func Before(fn func()) {
    testRun.before = fn
}

/*
    Naming function for describing the group of It() functions contained in the function argument.
*/
func Describe(title string, fn func()) {
    testRun.describe = title
    if testRun.before != nil {
        testRun.before()
    }
    fn()
    if testRun.after != nil {
        testRun.after()
    }
}

/*
    Called after each Describe() function call.
*/
func After(fn func()) {
    testRun.after = fn
}

/*
    Called before each It() function call.
*/
func BeforeEach(fn func()) {
    testRun.beforeEach = fn
}

/*
    Naming function for describing what should be return by the function argument.
*/
func It(title string, fn func()) {
    testRun.it = title
    if testRun.beforeEach != nil {
        testRun.beforeEach()
    }
    fn()
    if testRun.afterEach != nil {
        testRun.afterEach()
    }
}

/*
    Called after each It() function call.
*/
func AfterEach(fn func()) {
    testRun.afterEach = fn
}