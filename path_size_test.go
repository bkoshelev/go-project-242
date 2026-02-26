package code

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestGetSize(t *testing.T) {
	cases := []struct {
		path  string
		all   bool
		error error
		want  int64
	}{
		{"./testdata/512b.txt", false, nil, 512},
		{"./testdata/256k.txt", false, nil, 256 * KB},
		{"./testdata/1_m.txt", false, nil, MB},
		{"./testdata/unknown.txt", false, errors.New("test"), 0},
		{"./testdata", false, nil, 512 + (256 * KB) + MB},
		{"./testdata", true, nil, 512 + (256 * KB) + MB + (1000 * KB)},
		{"./testdata/empty_folder", false, nil, 0},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s", c.path)

		t.Run(name, func(t *testing.T) {
			got, e := GetSize(c.path, c.all)
			if c.error == nil && e != nil {
				t.Errorf("Не ожидали ошибку")
			}
			if c.error != nil && e == nil {
				t.Errorf("Ожидалась ошибка")
			}
			if got != c.want {
				t.Errorf("FormatSize(%s, %v) = %v, хотели %v", c.path, c.all, got, c.want)
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
