package util

import (
	"reflect"
	"testing"
)

func TestArrayToString(t *testing.T) {
	tests := []struct {
		in  []int
		out string
	}{
		{
			[]int{},
			"",
		},
		{
			[]int{1, 2, 3},
			"123",
		},
		{
			[]int{123},
			"123",
		},
		{
			[]int{0, 1, 2, 3},
			"0123",
		},
		{
			[]int{-1, 2, 3},
			"123",
		},
		{
			[]int{1, -2, 3},
			"123",
		},
	}

	for _, test := range tests {
		if ArrayToString(test.in) != test.out {
			t.Errorf("Expected: %s, Actual: %s", test.out, ArrayToString(test.in))
		}
	}
}

func TestArrayToInt(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 2, 3},
			123,
		},
		{
			[]int{0, 1, 2, 3},
			123,
		},
		{
			[]int{-1, 2, 3},
			123,
		},
		{
			[]int{-1, -2, 3},
			123,
		},
	}

	for _, test := range tests {
		if ArrayToInt(test.in) != test.out {
			t.Errorf("Expected: %d, Actual: %d", test.out, ArrayToInt(test.in))
		}
	}
}

func TestIntToArray(t *testing.T) {
	tests := []struct {
		n      int
		length int
		out    []int
	}{
		{
			0,
			1,
			[]int{0},
		},
		{
			123,
			5,
			[]int{0, 0, 1, 2, 3},
		},
		{
			123,
			-1,
			[]int{1, 2, 3},
		},
		{
			-123,
			-1,
			[]int{1, 2, 3},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(IntToArray(test.n, test.length), test.out) {
			t.Errorf("Expected: %v, Actual: %v", test.out, IntToArray(test.n, test.length))
		}
	}
}
