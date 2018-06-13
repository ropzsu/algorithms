package lmath

import "testing"

func Test_addStrings(t *testing.T) {
	r := addStrings("14435254636565634524133123124355245252", "97382327637821361813213232132132132134343898779")
	if r != "97382327652256616449778866656265255258699144031" {
		t.Fatal(r)
	}
}
