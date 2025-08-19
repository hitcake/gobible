package methods

import (
	"strings"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	Print(time.Hour)

	Print(new(strings.Replacer))
}
