package double_test

import (
	"fmt"
	"testing"

	"double.com"
)

func TestDouble2Returns4(t *testing.T) {
	t.Parallel()
	want := 4
	got := double.Double(2)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func ExampleDouble() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_two() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_three() {
	fmt.Println(double.Double(3))
	// Output:
	// 6
}
