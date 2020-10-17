package structfield

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name       string
		dst        interface{}
		src        interface{}
		wantErr    bool
		wantString string
	}{
		{
			"Test case only B",
			&struct {
				A string
				B string
			}{"-", "-"},
			struct {
				B string
				C string
			}{"B", "C"},
			false,
			"&{- B}",
		},
		{
			"Test case no pointer",
			struct {
				A string
				B string
			}{"-", "-"},
			struct {
				B string
				C string
			}{"B", "C"},
			true,
			"{- -}",
		},
		{
			"Test case other field type",
			&struct {
				A string
				B int
			}{"-", 5},
			struct {
				B string
				C string
			}{"B", "C"},
			false,
			"&{- 5}",
		},
		{
			"Tags",
			&struct {
				A string
				B string
				C string
			}{"-", "-", "-"},
			struct {
				A string `structfield:"nocopy"`
				B string
				C string
			}{"A", "B", "C"},
			false,
			"&{- B C}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.dst, tt.src); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr: %v", err, tt.wantErr)
			}
			if fmt.Sprintf("%v", tt.dst) != tt.wantString {
				t.Errorf("\nGot:  %v\nWant: %v", tt.dst, tt.wantString)
			}
		})
	}
}
