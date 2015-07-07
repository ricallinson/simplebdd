package simplebdd

import (
    "testing"
)

func TestRunner(t *testing.T) {

    Describe("Describe() and It()", func() {
        It("should return true after being called", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Skip()", func() {
        Skip("should return nothing", func() {
            AssertEqual(true, false)
        })
    })

    Report(t)
}
