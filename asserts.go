package simplebdd

/*
   Asserts that "a" and "b" are the same and stores the result for later use by the Report function.
*/
func AssertEqual(a interface{}, b interface{}) {

	test := false

	switch a.(type) {
	case string:
		test = a.(string) == b.(string)
	case bool:
		test = a.(bool) == b.(bool)
	case uint:
		test = a.(uint) == b.(uint)
	case uint8: // byte
		test = a.(uint8) == b.(uint8)
	case uint16:
		test = a.(uint16) == b.(uint16)
	case uint32:
		test = a.(uint32) == b.(uint32)
	case uint64:
		test = a.(uint64) == b.(uint64)
	case int:
		test = a.(int) == b.(int)
	case int8:
		test = a.(int8) == b.(int8)
	case int16:
		test = a.(int16) == b.(int16)
	case int32: // rune
		test = a.(int32) == b.(int32)
	case int64:
		test = a.(int64) == b.(int64)
	case float32:
		test = a.(float32) == b.(float32)
	case float64:
		test = a.(float64) == b.(float64)
	case complex64:
		test = a.(complex64) == b.(complex64)
	case complex128:
		test = a.(complex128) == b.(complex128)
	}

	record(test, a, b)
}

/*
   Asserts that "a" and "b" are different and stores the result for later use by the Report function.
*/
func AssertNotEqual(a interface{}, b interface{}) {

	test := false

	switch a.(type) {
	case string:
		test = a.(string) != b.(string)
	case bool:
		test = a.(bool) != b.(bool)
	case uint:
		test = a.(uint) != b.(uint)
	case uint8: // byte
		test = a.(uint8) != b.(uint8)
	case uint16:
		test = a.(uint16) != b.(uint16)
	case uint32:
		test = a.(uint32) != b.(uint32)
	case uint64:
		test = a.(uint64) != b.(uint64)
	case int:
		test = a.(int) != b.(int)
	case int8:
		test = a.(int8) != b.(int8)
	case int16:
		test = a.(int16) != b.(int16)
	case int32: // rune
		test = a.(int32) != b.(int32)
	case int64:
		test = a.(int64) != b.(int64)
	case float32:
		test = a.(float32) != b.(float32)
	case float64:
		test = a.(float64) != b.(float64)
	case complex64:
		test = a.(complex64) != b.(complex64)
	case complex128:
		test = a.(complex128) != b.(complex128)
	}

	record(test, a, b)
}
