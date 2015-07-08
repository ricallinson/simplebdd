# SimpleBDD Test Runner

[![Build Status](https://travis-ci.org/ricallinson/simplebdd.svg?branch=master)](https://travis-ci.org/ricallinson/simplebdd)

Super simple behavior-driven development style test writer for Go.

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
            It("should return that false is not true", func() {
                AssertNotEqual(false, true)
            })
            Skip("should return that 1 is not 2", func() {
                AssertNotEqual(1, 2)
            })
        })
        Report(t)
    }
