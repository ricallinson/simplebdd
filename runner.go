package simplebdd

import(
    "fmt"
)

/*
    result type
*/

type result struct {
    message string
    passed bool
}

/*
    TestRunner type.
*/

type TestRunner struct {
    describe string
    it string
    tests []result
}

/*
    Holds all tests run.
*/

var testRun = TestRunner{}

/*
    Store the result of an Assert call.
*/

func record(pass bool, got interface{}) {
    testRun.tests = append(testRun.tests, result{
        message: testRun.describe + " " + testRun.it + ", got " + fmt.Sprint(got),
        passed: pass,
    })
}

/*
    Grouping function for describing a function.
*/

func Describe(title string, fn func()) {
    testRun.describe = title
    fn()
}

/*
    Naming function for describing calling a function.
*/

func It(title string, fn func()) {
    testRun.it = title
    fn()
}