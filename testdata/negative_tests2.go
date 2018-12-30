package ntests2

// In this test suite, (2) option is always used. No warnings should be generated.

import (
	"strconv"
)

import (
	"errors"
)

import (
	"fmt"
)

var (
	_ = fmt.Printf
	_ = errors.New
	_ = strconv.Atoi
)

// T is an example type.
type T struct {
	integer int
}

func zeroValPtrAlloc() {
	_ = &T{}
	_ = &T{}
	_ = &T{}

	// Not a zero value allocation:
	_ = &T{integer: 1}
}

func emptySlice() {
	_ = []rune{}
	_ = []float64{}
	_ = []string{}
}

func emptyMap() {
	_ = map[rune]rune{}
	_ = map[*T]*T{}
	_ = map[int]int{}
}

func hexLit() {
	_ = 0xFF
	_ = 0xABCDEF
	_ = 0xABCD
}

func rangeCheck(x, low, high int) {
	_ = low <= x && x <= high
	_ = low < x+1 || x+1 < high
	_ = low < x && x < high
}

func andNot(x, y int) {
	_ = x & ^y
	_ = 123 & ^x
	_ = (x + 100) & ^(y + 2)
}

func floatLit() {
	_ = .43
	_ = 1.
	_ = .0
}

func labelCase() {
AllUpper:
Foo:
UpperCamelCase:
LowerCamelCase:
	goto AllUpper
	goto Foo
	goto UpperCamelCase
	goto LowerCamelCase
}

func untypedConstCoerce() {
	const zero = 0

	var _ = int(zero)
	var _ = int32(10)
	var _ = int64(zero + 1)
}

func threeArgs(a, b, c int) {}

func argListParens() {
	threeArgs(
		1,
		2,
		3,
	)
	threeArgs(1,
		2,
		3,
	)
	threeArgs(
		1,
		2,
		3,
	)
}

func nonZeroLenTestChecker() {
	var (
		s  string
		b  []byte
		m  map[int]int
		ch chan int
	)

	// Strings are ignored.
	_ = len(s) >= 1
	_ = len(s) >= 1
	_ = len(s) >= 1

	_ = len(b) > 0
	_ = len(m) > 0
	_ = len(ch) > 0
	_ = len(ch) > 0
}

func defaultCaseOrder(x int, v interface{}) {
	switch x {
	case 10:
	default:
	}

	switch v.(type) {
	case int:
	case string:
	default:
	}

	switch {
	case x > 20:
	default:
	}
}
