# SimpleBDD Test Runner

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
        })
        Describe("AssertNotEqual()", func() {
            It("should return that false is not true", func() {
                AssertNotEqual(false, true)
            })
        })
        Report(t)
    }
