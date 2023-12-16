package main

import "testing"

func TestMutateWhenHasCapacity(t *testing.T) {
	toAppend := 123
	sliceLastIndex := 3
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[:sliceLastIndex]
	slice = Append(slice, toAppend)

	if len(slice) != sliceLastIndex+1 {
		t.Fatalf("Expected slice with length 4, got %d", len(slice))
	}

	if slice[sliceLastIndex] != toAppend {
		t.Fatalf("Expected to append element %v to the slice, got %v", toAppend, slice[3])
	}

	if arr[sliceLastIndex] != toAppend {
		t.Fatalf("Expected arrays element with index 3 to be %v, got %v", toAppend, arr[3])
	}
}

func TestNotMutateWhenDoesNotHaveCapacity(t *testing.T) {
	toAppend := 123
	arr := [6]int{1, 2, 3, 4, 5, 6}
	oldLen := len(arr)
	slice := arr[:]
	slice = Append(slice, toAppend)

	if len(slice) != oldLen+1 {
		t.Fatalf("Expected slice with length 7, got %d", len(slice))
	}

	if slice[oldLen] != toAppend {
		t.Fatalf("Expected to append element %v to the slice, got %v", toAppend, slice[3])
	}

	if arr[oldLen-1] != 6 {
		t.Fatal("Expected to not mutate array")
	}
}

func TestMutateWhenDoesNotHaveCapacity(t *testing.T) {
	toAppend := 123
	sliceFirstIndex := 3
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[sliceFirstIndex:]
	slice = Append(slice, toAppend)

	if len(slice) != len(arr)-sliceFirstIndex+1 {
		t.Fatalf("Expected slice with length 4, got %d", len(slice))
	}

	if slice[len(slice)-1] != toAppend {
		t.Fatalf("Expected to append element %v to the slice, got %v", toAppend, slice[3])
	}

	if arr[len(arr)-1] == toAppend {
		t.Fatalf("Expected arrays element with index 5 to be %v, got %v", 6, arr[5])
	}
}

func TestAppendVariadicArgs(t *testing.T) {
	args := []int{1, 2, 3}
	s := make([]int, 0, 3)
	s = Append(s, args...)

	if len(s) != len(args) {
		t.Fatalf("expected to have %d arguments, got %d", len(args), len(s))
	}

	for idx, expected := range args {
		if s[idx] != expected {
			t.Errorf("Expected %v, actual: %v", expected, s[idx])
		}
	}
}

func TestAppendNilSlice(t *testing.T) {
	args := []int{1, 2, 3}
	var s []int = nil
	s = Append(s, args...)

	if len(s) != len(args) {
		t.Fatalf("expected to have %d arguments, got %d", len(args), len(s))
	}

	for idx, expected := range args {
		if s[idx] != expected {
			t.Errorf("Expected %v, actual: %v", expected, s[idx])
		}
	}
}
