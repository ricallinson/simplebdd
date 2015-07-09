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

## Xunit Reporting

If you want to output Junit style reports, call `Xunit()` just prior to `Report()` passing the given `testing.T` variable. This will create a folder named _test-results_ that contains __xml__ files for each executed `Describe` block.

    func TestSomething(t *testing.T) {
        Describe("AssertEqual()", func() {
            It("should return that true is true", func() {
                AssertEqual(true, true)
            })
        })
        Xunit(t)
        Report(t)
    }

## Testing

    go test -cover

## Coverage
    
    go test -coverprofile=coverage.out; go tool cover -html=coverage.out -o=index.html
