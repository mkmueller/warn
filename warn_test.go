// Copyright (c) 2018 Mark K Mueller <github.com/mkmueller>

package warn

import (
	"fmt"
	"testing"
)

func TestNew (t *testing.T) {
	warn := New("This is a warning")
	warn.Append("This is a another warning")
	if warn.Warnings()[0] != "This is a warning" {
		t.Fail()
	}
	if warn.Warnings()[1] != "This is a another warning" {
		t.Fail()
	}
}

func ExampleNew () {
	warn := New("This is a warning")
	warn.Append("This is a another warning")
	for _,w := range warn.Warnings() {
		fmt.Printf("%s\n", w)
	}
	// Output: This is a warning
	// This is a another warning
}
