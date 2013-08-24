package simplebdd

import(
    "fmt"
    "testing"
)

/*
    Really simple report.
*/

func Report(t *testing.T) {

    ok := "✓"
    err := "✖"

    failed := 0
    passed := 0

    colorNormal := "\x1b[0m"
    colorTitled := "\x1b[37m"
    colorReport := "\x1b[90m"
    colorPassed := "\x1b[92m"
    colorFailed := "\x1b[31m"

    fmt.Println("\x1b[90m")

    group := ""

    for _, test := range testRun.tests {

        if group != test.group {
            group = test.group
            fmt.Println(colorTitled + group + colorReport + "\n")
        }

        if test.passed {
            fmt.Println("\t" + colorPassed + ok + colorReport + " " + test.message)
            passed++
        } else {
            t.Fail()
            msg := colorFailed
            msg += err + " " + test.message + ", got [" + test.result + "]"
            msg += colorReport
            fmt.Println("\t" + msg)
            failed++
        }
    }

    status := colorFailed + err + colorReport

    if failed == 0 {
        status = colorPassed + ok + colorReport
    }

    fmt.Println()
    fmt.Println(status, len(testRun.tests), "tests complete")

    fmt.Println(colorNormal)

    /*
        Once reported clean the test runner for the next run.
    */

    testRun = TestRunner{}
}