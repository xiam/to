package to

import (
	"fmt"
	"testing"
)

func TestNumeric(t *testing.T) {
	if String(0) != "0" {
		t.Fatalf("Test failed.")
	}
	if String(-0) != "0" {
		t.Fatalf("Test failed.")
	}
	if String(1) != "1" {
		t.Fatalf("Test failed.")
	}
	if String(-1) != "-1" {
		t.Fatalf("Test failed.")
	}
	if String(10) != "10" {
		t.Fatalf("Test failed.")
	}
	if String(-10) != "-10" {
		t.Fatalf("Test failed.")
	}
	if String(int64(9223372036854775807)) != "9223372036854775807" {
		t.Fatalf("Test failed.")
	}
	if String(int64(-9223372036854775807)) != "-9223372036854775807" {
		t.Fatalf("Test failed.")
	}
	if String(uint64(18446744073709551615)) != "18446744073709551615" {
		t.Fatalf("Test failed.")
	}

	fmt.Printf("%v\n", 0.1)
	fmt.Printf("%v\n", 0.1123)
	fmt.Printf("%v\n", -0.1235)
	fmt.Printf("%v\n", 0.1765428833765552677777766)
	fmt.Printf("%v\n", 99999999999999.1765428833765552677777766)
}
