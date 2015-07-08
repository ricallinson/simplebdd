package simplebdd

import (
    "testing"
)

func TestAsserts(t *testing.T) {

    Describe("AssertEqual()", func() {
        It("should return string true", func() {
            AssertEqual("t", "t")
        })
    })

    Describe("AssertNotEqual()", func() {
        It("should return string true", func() {
            AssertNotEqual("t", "f")
        })
    })

    Report(t)
}
