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
    Holds all tests run.
*/

var testRun = testRunner{}

/*
    Store the result of an Assert call.
*/

func record(pass bool, got interface{}) {
    testRun.tests = append(testRun.tests, result{
        group: testRun.describe,
        message: testRun.it,
        result: fmt.Sprint(got),
        passed: pass,
    })
}

/*
    Function to call before each It().
*/

func Before(fn func()) {
    testRun.before = fn
}

/*
    Grouping function for describing a function.
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
    Function to call after each It().
*/

func After(fn func()) {
    testRun.after = fn
}

/*
    Function to call before each It().
*/

func BeforeEach(fn func()) {
    testRun.beforeEach = fn
}

/*
    Naming function for describing calling a function.
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
    Function to call after each It().
*/

func AfterEach(fn func()) {
    testRun.afterEach = fn
}