package code

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

const (
	B_512                       = 512
	KB_256                      = 256 * KB
	KB_1000                     = 1000 * KB
	TEST_DATA_SIZE              = B_512 + KB_256 + MB
	TEST_DATA_SIZE_ALL          = B_512 + KB_256 + MB + KB_1000
	INTERNAL_TEST_DATA_SIZE     = KB_256
	INTERNAL_TEST_DATA_SIZE_ALL = KB_256 + KB_1000
)

func TestGetSize(t *testing.T) {
	cases := []struct {
		path      string
		all       bool
		recursive bool
		error     error
		want      int64
	}{
		{"./testdata/512b.txt", false, false, nil, B_512},
		{"./testdata/256k.txt", false, false, nil, KB_256},
		{"./testdata/1_m.txt", false, false, nil, MB},
		{"./testdata/unknown.txt", false, false, errors.New("test"), 0},
		{"./testdata", false, false, nil, TEST_DATA_SIZE},
		{"./testdata", true, false, nil, TEST_DATA_SIZE_ALL},
		{"./testdata/empty_folder", false, false, nil, 0},
		{"./testdata", false, true, nil, TEST_DATA_SIZE + INTERNAL_TEST_DATA_SIZE},
		{"./testdata", true, true, nil, TEST_DATA_SIZE_ALL + INTERNAL_TEST_DATA_SIZE_ALL},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s_%v_%v", c.path, c.all, c.recursive)

		t.Run(name, func(t *testing.T) {
			got, e := GetSize(c.path, c.all, c.recursive)
			if c.error == nil && e != nil {
				t.Errorf("Не ожидали ошибку %q", c.error)
			}
			if c.error != nil && e == nil {
				t.Errorf("Ожидалась ошибка")
			}
			if got != c.want {
				t.Errorf("FormatSize(%s, %v, %v) = %v, хотели %v", c.path, c.all, c.recursive, got, c.want)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	cases := []struct {
		size  int64
		human bool
		want  string
	}{
		{0, false, "0B"},
		{-1, false, "0B"},
		{2, false, "2B"},
		{1024, false, "1024B"},
		{0, true, "0B"},
		{1023, true, "1023B"},
		{KB, true, "1.0KB"},
		{KB * 2, true, "2.0KB"},
		{MB, true, "1.0MB"},
		{GB, true, "1.0GB"},
		{TB, true, "1.0TB"},
		{PB, true, "1.0PB"},
		{EB, true, "1.0EB"},
		{math.MaxInt, true, "8.0EB"},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%d_%v", c.size, c.human)

		t.Run(name, func(t *testing.T) {
			got := FormatSize(c.size, c.human)
			if got != c.want {
				t.Errorf("FormatSize(%d, %v) = %v, хотели %v", c.size, c.human, got, c.want)
			}
		})
	}
}
