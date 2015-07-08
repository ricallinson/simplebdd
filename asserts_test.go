package simplebdd

import (
    "testing"
)

func TestAsserts(t *testing.T) {

    Describe("AssertEqual()", func() {
        It("should return string true", func() {
            AssertEqual("t", "t")
        })
        It("should return bool true", func() {
            AssertEqual(true, true)
        })
        It("should return uint true", func() {
            AssertEqual(uint(1), uint(1))
        })
        It("should return uint8 true", func() {
            AssertEqual(uint8(1), uint8(1))
        })
        It("should return uint16 true", func() {
            AssertEqual(uint16(1), uint16(1))
        })
        It("should return uint32 true", func() {
            AssertEqual(uint32(1), uint32(1))
        })
        It("should return uint64 true", func() {
            AssertEqual(uint64(1), uint64(1))
        })
        It("should return int true", func() {
            AssertEqual(int(-1), int(-1))
        })
        It("should return int8 true", func() {
            AssertEqual(int8(-1), int8(-1))
        })
        It("should return int16 true", func() {
            AssertEqual(int16(-1), int16(-1))
        })
        It("should return int32 true", func() {
            AssertEqual(int32(-1), int32(-1))
        })
        It("should return int64 true", func() {
            AssertEqual(int64(-1), int64(-1))
        })
        It("should return float32 true", func() {
            AssertEqual(float32(1.1), float32(1.1))
        })
        It("should return float64 true", func() {
            AssertEqual(float64(1.1), float64(1.1))
        })
        It("should return complex64 true", func() {
            AssertEqual(complex64(2.1), complex64(2.1))
        })
        It("should return complex128 true", func() {
            AssertEqual(complex128(2.1), complex128(2.1))
        })
    })

    Describe("AssertNotEqual()", func() {
        It("should return string true", func() {
            AssertNotEqual("t", "f")
        })
        It("should return bool true", func() {
            AssertNotEqual(true, false)
        })
        It("should return uint true", func() {
            AssertNotEqual(uint(1), uint(2))
        })
        It("should return uint8 true", func() {
            AssertNotEqual(uint8(1), uint8(2))
        })
        It("should return uint16 true", func() {
            AssertNotEqual(uint16(1), uint16(2))
        })
        It("should return uint32 true", func() {
            AssertNotEqual(uint32(1), uint32(2))
        })
        It("should return uint64 true", func() {
            AssertNotEqual(uint64(1), uint64(2))
        })
        It("should return int true", func() {
            AssertNotEqual(int(-1), int(-2))
        })
        It("should return int8 true", func() {
            AssertNotEqual(int8(-1), int8(-2))
        })
        It("should return int16 true", func() {
            AssertNotEqual(int16(-1), int16(-2))
        })
        It("should return int32 true", func() {
            AssertNotEqual(int32(-1), int32(-2))
        })
        It("should return int64 true", func() {
            AssertNotEqual(int64(-1), int64(-2))
        })
        It("should return float32 true", func() {
            AssertNotEqual(float32(1.1), float32(1.2))
        })
        It("should return float64 true", func() {
            AssertNotEqual(float64(1.1), float64(1.2))
        })
        It("should return complex64 true", func() {
            AssertNotEqual(complex64(2.1), complex64(2.2))
        })
        It("should return complex128 true", func() {
            AssertNotEqual(complex128(2.1), complex128(2.2))
        })
    })

    Report(t)
}
