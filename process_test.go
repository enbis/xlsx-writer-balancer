package wb

import (
	"testing"
)

func TestMultiProcess(t *testing.T) {

	Init()

	var a = []int{}
	var b = []int{}

	for i := 0; i < 1000; i++ {
		a = append(a, i)
		b = append(b, i+1000)
	}

	tests := []struct {
		namefile string
		values   []int
	}{
		{namefile: "one", values: a},
		{namefile: "two", values: b},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.namefile, func(t *testing.T) {
			t.Parallel()
			for _, v := range tc.values {
				Process(tc.namefile, v)
			}
		})
	}
}

func TestProcessFileExists(t *testing.T) {
	Init()

	var a = []int{}

	for i := 3000; i < 4000; i++ {
		a = append(a, i)
	}

	tests := []struct {
		namefile string
		values   []int
	}{
		{namefile: "one", values: a},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.namefile, func(t *testing.T) {
			t.Parallel()
			for _, v := range tc.values {
				Process(tc.namefile, v)
			}
		})
	}
}
