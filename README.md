# SimpleBDD Test Runner

Super simple BDD style test runner for Go.

    package something

    import(
        "testing"
        . "github.com/ricallinson/simplebdd"
    )

    func TestSomething(t *testing.T) {
        Describe("AssertEqual()", func() {
            It("should return true", func() {
                AssertEqual(true, true)
            })
        })
        Report(t)
    }
