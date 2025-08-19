package format

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatAtom(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // 1
	fmt.Println(Any(d))                  // 1
	fmt.Println(Any([]int64{x}))         // []int64 0xc000010398
	fmt.Println(Any([]time.Duration{d})) // []time.Duration 0xc0000103b0
}
