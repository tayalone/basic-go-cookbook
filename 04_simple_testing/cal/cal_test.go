package cal

import "testing"

func TestSumOfFirstThree(t *testing.T) {
	given := 3
	want := 6

	get := SumOfFirst(given)
	if want != get {
		t.Errorf("given %d but get %d\n", given, get)
	}
}
