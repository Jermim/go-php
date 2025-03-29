package main

import (
	"fmt"
	"testing"
)

func TestPhpArray_all(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	if tru := Array_all(numbers, func(n int, i int) bool {
		return i+1 == n
	}); !tru {
		t.Logf("expected 'true' but got 'false'")
		t.FailNow()
	}

	if tru := Array_all(numbers, func(n int, i int) bool {
		if i == 1 {
			return i == n
		}

		return i+1 == n
	}); tru {
		t.Logf("expected 'false' but got 'true'")
		t.FailNow()
	}

}

func TestPhpArray_any(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	if tru := Array_any(numbers, func(n int, i int) bool {
		if i == 1 {
			return n == 2
		}
		return false
	}); !tru {
		t.Logf("expected 'false' but got 'true'")
		t.FailNow()
	}

	if tru := Array_any(numbers, func(n int, i int) bool {
		return false
	}); tru {
		t.Logf("expected 'true' but got 'false'")
		t.FailNow()
	}
}

func TestPhpArray_chunk(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	result := Array_chunk(numbers, 2)

	fmt.Printf("%+v\n", result)

	numbers = []int{1}

	result = Array_chunk(numbers, 2)

	fmt.Printf("%+v\n", result)

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}

	result = Array_chunk(numbers, 9)

	fmt.Printf("%+v\n", result)
}

func TestPhpArray_fill(t *testing.T) {

	arr := Array_fill(3, 3, "hello")

	fmt.Printf("%+v : %d\n", arr, len(arr))
}

func TestPhpArray_pop(t *testing.T) {

	val := Array_Pop(&[]int{1, 2, 3, 4, 5})

	if expected := 5; expected != val {
		t.Logf("expected '%d' got '%d'", expected, val)
		t.FailNow()
	}
}

func TestPhpArray_map(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	result := Array_map(numbers, func(n int) string {
		return fmt.Sprintf("the number is %d", n)
	})

	for i, s := range result {
		if expected := fmt.Sprintf("the number is %d", numbers[i]); expected != s {
			t.Logf("expected '%s' but got '%s'", expected, s)
			t.FailNow()
		}
	}
}
