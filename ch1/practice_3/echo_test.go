package practice_3

import (
	"strings"
	"testing"
)

/*
 */
var test = strings.Split(`This chapter documents arrays, a fundamental datatype in JavaScript and in most other programming languages. 
An array is an ordered collection of values. Each value is called an element, and each element has a numeric position in the array, known as its index. 
JavaScript arrays are untyped: an array element may be of any type, and different elements of the same array may be of different types. `, " ")

func BenchmarkEchoV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoV1(test)
	}
}

func BenchmarkEchoV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoV2(test)
	}
}

/*
goos: darwin
goarch: amd64
pkg: gobible/ch1/practice_3
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkEcho_v1
BenchmarkEcho_v1-8   	  182455	      6387 ns/op
BenchmarkEcho_v2
BenchmarkEcho_v2-8   	 1366479	       814.0 ns/op
PASS
*/
