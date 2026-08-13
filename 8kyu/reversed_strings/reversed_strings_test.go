package reversed_strings

import "testing"

func TestReverse(t *testing.T) {
	result := Reverse("he")

	expect := "eh"
	if result != expect {
		t.Error("Expected", expect, "got", result)
	}
}
