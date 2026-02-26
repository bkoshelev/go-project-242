package code

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetSize(t *testing.T) {
	cases := []struct {
		path  string
		error error
		want  int64
	}{
		{"./testdata/512b.txt", nil, 512},
		{"./testdata/256k.txt", nil, 256 * KB},
		{"./testdata/1_m.txt", nil, MB},
		{"./testdata/unknown.txt", errors.New("test"), 0},
		{"./testdata", nil, 512 + (256 * KB) + MB},
		{"./testdata/empty_folder", nil, 0},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s", c.path)

		t.Run(name, func(t *testing.T) {
			got, e := GetSize(c.path)
			if c.error == nil && e != nil {
				t.Errorf("Не ожидали ошибку")
			}
			if c.error != nil && e == nil {
				t.Errorf("Ожидалась ошибка")
			}
			if got != c.want {
				t.Errorf("FormatSize(%s) = %v, хотели %v", c.path, got, c.want)
			}
		})
	}
}
