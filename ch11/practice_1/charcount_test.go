package practice_1

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
	This is a new world.
	Enjoy int.   
    我在中国
    `))
	countResult, err := CountCharacters(reader)
	if err != nil {
		t.Fatal(err)
	}
	expectedCount := map[rune]int{
		' ':  16,
		'w':  2,
		'o':  2,
		'l':  1,
		'd':  1,
		'E':  1,
		'n':  3,
		'e':  1,
		'.':  2,
		'j':  1,
		'y':  1,
		'我':  1,
		'T':  1,
		'i':  3,
		'r':  1,
		'在':  1,
		'\n': 4,
		'\t': 2,
		'a':  1,
		't':  1,
		'中':  1,
		'国':  1,
		'h':  1,
		's':  2,
	}
	if !reflect.DeepEqual(expectedCount, countResult.Counts) {
		t.Fatalf("counts are different")
	}
	fmt.Printf("rune\tcount\n")
	if countResult.Utflen[1] != 47 {
		t.Error("counts.Utflen[1]!=47")
	}
	if countResult.Utflen[3] != 4 {
		t.Error("counts.Utflen[3]!=4")
	}
	if countResult.Invalid > 0 {
		t.Errorf("invalid %d", countResult.Invalid)
	}
}
