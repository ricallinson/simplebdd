package simplebdd

import(
    "fmt"
    "testing"
)

// Prints to stdout a list of the tests and if they passed or failed.
func Report(t *testing.T) {

    ok := "✓"
    err := "✖"

    failed := 0
    passed := 0
    skipped := 0

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
            fmt.Println("\n" + colorTitled + group + colorReport + "\n")
        }

        if test.skipped {
            fmt.Println("\t" + "  Skipped: " + test.message)
            skipped++
        } else if test.passed {
            fmt.Println("\t" + colorPassed + ok + colorReport + " " + test.message)
            passed++
        } else {
            t.Fail()
            msg := colorFailed
            msg += err + " " + test.message + ", expected [" + test.expect + "], got [" +test.result + "]"
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
    fmt.Println(status, len(testRun.tests) - skipped, "tests complete")

    if skipped > 0 {
        fmt.Println(" ", skipped, "tests skipped")
    }

    fmt.Println(colorNormal)

    // Once reported clean the test runner for the next run.
    testRun = testRunner{}
}