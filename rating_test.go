package advent

import "testing"

func Test_debug(t *testing.T) {
	SetDebug(t)
	defer SetDebug(nil)
	LifeSupportRating("testdata/3.input", 12)
}
