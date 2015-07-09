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
}

func TestRunnerBefore(t *testing.T) {

	test := false

	Before(func() {
		test = true
	})

	Describe("Before()", func() {
		It("should return true becuase test is true", func() {
			AssertEqual(test, true)
			test = false
		})
	})

	Describe("Before() again", func() {
		It("should return true becuase test is true", func() {
			AssertEqual(test, true)
		})
	})

	Report(t)
}

func TestRunnerAfter(t *testing.T) {

	test := false

	After(func() {
		test = false
	})

	Describe("After()", func() {
		It("should return true becuase test is false", func() {
			AssertEqual(test, false)
			test = true
		})
	})

	Describe("After() again", func() {
		It("should return true becuase test is false", func() {
			AssertEqual(test, false)
		})
	})

	Report(t)
}

func TestRunnerBeforeEach(t *testing.T) {

	Describe("BeforeEach()", func() {

		test := false

		BeforeEach(func() {
			test = true
		})

		It("should return true becuase test is true", func() {
			AssertEqual(test, true)
			test = false
		})

		It("should return true becuase test is still true", func() {
			AssertEqual(test, true)
		})
	})

	Report(t)
}

func TestRunnerAfterEach(t *testing.T) {

	Describe("AfterEach()", func() {

		test := true

		AfterEach(func() {
			test = false
		})

		It("should return true becuase test is true", func() {
			AssertEqual(test, true)
			test = true
		})

		It("should return true becuase test is still false", func() {
			AssertEqual(test, false)
		})
	})

	Report(t)
}
